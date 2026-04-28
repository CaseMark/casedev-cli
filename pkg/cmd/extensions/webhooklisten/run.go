package webhooklisten

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/urfave/cli/v3"
)

const defaultDescription = "casedev webhook listen (ephemeral ngrok tunnel)"

type endpointCreateRequest struct {
	URL              string   `json:"url"`
	Description      string   `json:"description,omitempty"`
	EventTypeFilters []string `json:"eventTypeFilters"`
}

type endpointCreateResponse struct {
	ID            string   `json:"id"`
	SigningSecret string   `json:"signingSecret"`
	URL           string   `json:"url"`
	Description   string   `json:"description"`
	EventTypes    []string `json:"eventTypeFilters"`
}

type ngrokTunnelListResponse struct {
	Tunnels []struct {
		Name      string `json:"name"`
		PublicURL string `json:"public_url"`
		Proto     string `json:"proto"`
	} `json:"tunnels"`
}

type runtime struct {
	printMode string
	forwardTo string
	stderr    io.Writer
	http      *http.Client
}

type server struct {
	listener net.Listener
	server   *http.Server
	runtime  runtime
}

func handleRun(ctx context.Context, cmd *cli.Command) error {
	if _, err := exec.LookPath("ngrok"); err != nil {
		return fmt.Errorf("ngrok not found in PATH. Install it first (for example: brew install ngrok)")
	}

	cfg, err := loadConfig(false)
	if err != nil && !errors.Is(err, errConfigNotFound) {
		return err
	}

	events := cfg.Events
	if raw := strings.TrimSpace(cmd.String("events")); raw != "" {
		events = normalizePatterns(strings.Split(raw, ","))
	}
	if len(events) == 0 {
		events = []string{"*"}
	}

	forwardTo := cfg.ForwardTo
	if cmd.IsSet("forward-to") {
		forwardTo = strings.TrimSpace(cmd.String("forward-to"))
	}
	printMode := strings.ToLower(strings.TrimSpace(cmd.String("print")))

	listener, port, err := startServer(cmd.Int("tunnel-port"), printMode, forwardTo)
	if err != nil {
		return err
	}
	defer listener.close()

	ngrokCmd, err := startNgrokProcess(port)
	if err != nil {
		return err
	}
	defer stopNgrokProcess(ngrokCmd)

	tunnelURL, err := waitForNgrokTunnelURL(ctx)
	if err != nil {
		return err
	}

	endpoint, err := createEndpoint(ctx, cmd, tunnelURL, events)
	if err != nil {
		return err
	}

	fmt.Fprintf(cmd.ErrWriter, "✓ Tunnel active: %s\n", tunnelURL)
	fmt.Fprintf(cmd.ErrWriter, "✓ Registered webhook endpoint (%s) for: %s\n", endpoint.ID, strings.Join(events, ", "))
	if forwardTo != "" {
		fmt.Fprintf(cmd.ErrWriter, "  Forwarding to: %s\n", forwardTo)
	}
	if endpoint.SigningSecret != "" {
		fmt.Fprintf(cmd.ErrWriter, "  Signing secret: %s\n", endpoint.SigningSecret)
	}
	fmt.Fprintln(cmd.ErrWriter)
	fmt.Fprintln(cmd.ErrWriter, "  Ready. Webhooks delivered via the full production pipeline.")
	fmt.Fprintln(cmd.ErrWriter, "  Press Ctrl-C to stop and clean up.")

	sigCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-sigCtx.Done()

	cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := deleteEndpoint(cleanupCtx, cmd, endpoint.ID); err != nil {
		fmt.Fprintf(cmd.ErrWriter, "warning: failed to delete webhook endpoint %s: %v\n", endpoint.ID, err)
	} else {
		fmt.Fprintf(cmd.ErrWriter, "✓ Deleted webhook endpoint %s\n", endpoint.ID)
	}
	fmt.Fprintln(cmd.ErrWriter, "✓ Tunnel closed")
	return nil
}

func startServer(port int, printMode string, forwardTo string) (*server, int, error) {
	addr := "127.0.0.1:0"
	if port > 0 {
		addr = fmt.Sprintf("127.0.0.1:%d", port)
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, 0, err
	}
	port = ln.Addr().(*net.TCPAddr).Port
	srv := &server{
		listener: ln,
		runtime: runtime{
			printMode: printMode,
			forwardTo: forwardTo,
			stderr:    os.Stderr,
			http:      &http.Client{Timeout: 15 * time.Second},
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", srv.handleWebhook)
	srv.server = &http.Server{Handler: mux}
	go func() {
		_ = srv.server.Serve(ln)
	}()
	return srv, port, nil
}

func (s *server) close() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = s.server.Shutdown(ctx)
}

func (s *server) handleWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	_ = r.Body.Close()

	printEvent(s.runtime, body)

	statusLine := ""
	if s.runtime.forwardTo != "" {
		statusLine = forwardRequest(s.runtime, r, body)
	}
	if statusLine != "" && s.runtime.printMode == "pretty" {
		fmt.Fprintln(s.runtime.stderr, statusLine)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func printEvent(rt runtime, body []byte) {
	if rt.printMode == "silent" {
		return
	}
	if rt.printMode == "json" {
		fmt.Fprintln(rt.stderr, string(body))
		return
	}

	var payload struct {
		Type       string `json:"type"`
		OccurredAt string `json:"occurred_at"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Fprintf(rt.stderr, "%s\n", string(body))
		return
	}
	if payload.Type == "" {
		fmt.Fprintf(rt.stderr, "%s\n", string(body))
		return
	}
	if payload.OccurredAt != "" {
		fmt.Fprintf(rt.stderr, "%s → %s\n", payload.OccurredAt, payload.Type)
		return
	}
	fmt.Fprintf(rt.stderr, "%s\n", payload.Type)
}

func forwardRequest(rt runtime, incoming *http.Request, body []byte) string {
	req, err := http.NewRequestWithContext(incoming.Context(), incoming.Method, rt.forwardTo, bytes.NewReader(body))
	if err != nil {
		return fmt.Sprintf("warning: forward failed: %v", err)
	}
	for key, values := range incoming.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	resp, err := rt.http.Do(req)
	if err != nil {
		return fmt.Sprintf("warning: forward failed: %v", err)
	}
	defer resp.Body.Close()
	return fmt.Sprintf("→ forwarded (%s)", resp.Status)
}

func startNgrokProcess(port int) (*exec.Cmd, error) {
	cmd := exec.Command("ngrok", "http", fmt.Sprintf("%d", port), "--log", "stdout", "--log-format", "json")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}

func stopNgrokProcess(cmd *exec.Cmd) {
	if cmd == nil || cmd.Process == nil {
		return
	}
	_ = cmd.Process.Signal(syscall.SIGTERM)
	_, _ = cmd.Process.Wait()
}

func waitForNgrokTunnelURL(ctx context.Context) (string, error) {
	client := &http.Client{Timeout: 2 * time.Second}
	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}
		resp, err := client.Get("http://127.0.0.1:4040/api/tunnels")
		if err == nil {
			var payload ngrokTunnelListResponse
			func() {
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					return
				}
				if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
					return
				}
			}()
			for _, tunnel := range payload.Tunnels {
				if strings.HasPrefix(tunnel.PublicURL, "https://") {
					return tunnel.PublicURL, nil
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	return "", errors.New("timed out waiting for ngrok tunnel URL")
}

func createEndpoint(ctx context.Context, cmd *cli.Command, tunnelURL string, events []string) (endpointCreateResponse, error) {
	payload := endpointCreateRequest{
		URL:              tunnelURL,
		Description:      defaultDescription,
		EventTypeFilters: events,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return endpointCreateResponse{}, err
	}
	client := &http.Client{Timeout: 15 * time.Second}
	baseURL := strings.TrimRight(cmd.String("base-url"), "/")
	if baseURL == "" {
		baseURL = "https://api.case.dev"
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/webhooks/v1/endpoints", bytes.NewReader(body))
	if err != nil {
		return endpointCreateResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if apiKey := cmd.String("api-key"); apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return endpointCreateResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		return endpointCreateResponse{}, fmt.Errorf("create webhook endpoint: %s: %s", resp.Status, strings.TrimSpace(string(data)))
	}
	var endpoint endpointCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&endpoint); err != nil {
		return endpointCreateResponse{}, err
	}
	return endpoint, nil
}

func deleteEndpoint(ctx context.Context, cmd *cli.Command, endpointID string) error {
	client := &http.Client{Timeout: 15 * time.Second}
	baseURL := strings.TrimRight(cmd.String("base-url"), "/")
	if baseURL == "" {
		baseURL = "https://api.case.dev"
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, baseURL+"/webhooks/v1/endpoints/"+endpointID, nil)
	if err != nil {
		return err
	}
	if apiKey := cmd.String("api-key"); apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete webhook endpoint: %s: %s", resp.Status, strings.TrimSpace(string(data)))
	}
	return nil
}
