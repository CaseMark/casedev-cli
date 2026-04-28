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
	ID            string
	SigningSecret string
}

// parseCreateEndpointResponse extracts the endpoint id and signing
// secret from the API response. It tolerates the two response shapes
// the casedev API has historically used — fields at the top level, or
// wrapped inside a "data" object — and accepts either "id" or
// "endpointId" as the identifier key. The webhook delivery payloads
// already use the latter (data.endpointId), so the create response
// likely follows the same convention.
//
// Returns an error containing the raw body when no id can be located,
// so an unfamiliar response shape surfaces in the user-visible error
// rather than silently producing a broken cleanup URL.
func parseCreateEndpointResponse(body []byte) (endpointCreateResponse, error) {
	var raw map[string]any
	if err := json.Unmarshal(body, &raw); err != nil {
		return endpointCreateResponse{}, fmt.Errorf("decode create response: %w (body: %s)", err, strings.TrimSpace(string(body)))
	}

	sources := []map[string]any{raw}
	for _, wrapper := range []string{"data", "endpoint", "result", "payload"} {
		if nested, ok := raw[wrapper].(map[string]any); ok {
			sources = append(sources, nested)
		}
	}

	pick := func(keys ...string) string {
		for _, src := range sources {
			for _, key := range keys {
				if v, ok := src[key].(string); ok && v != "" {
					return v
				}
			}
		}
		return ""
	}

	out := endpointCreateResponse{
		ID:            pick("id", "endpointId"),
		SigningSecret: pick("signingSecret", "secret"),
	}
	if out.ID == "" {
		return out, fmt.Errorf("create webhook endpoint: response missing id field (body: %s)", strings.TrimSpace(string(body)))
	}
	return out, nil
}

type ngrokTunnelListResponse struct {
	Tunnels []struct {
		PublicURL string `json:"public_url"`
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

	cleanupOrphanEndpoints(ctx, cmd)

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
	if endpoint.SigningSecret != "" && cmd.Bool("show-secret") {
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
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return endpointCreateResponse{}, fmt.Errorf("read create response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return endpointCreateResponse{}, fmt.Errorf("create webhook endpoint: %s: %s", resp.Status, strings.TrimSpace(string(respBody)))
	}
	return parseCreateEndpointResponse(respBody)
}

func deleteEndpoint(ctx context.Context, cmd *cli.Command, endpointID string) error {
	if endpointID = strings.TrimSpace(endpointID); endpointID == "" {
		return fmt.Errorf("delete webhook endpoint: missing endpoint id (cleanup skipped)")
	}
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

// cleanupOrphanEndpoints lists existing webhook endpoints and removes
// any whose description matches the marker this tool writes. This
// undoes accumulated orphans from prior runs that exited without a
// successful DELETE (SIGKILL, network failure during cleanup, the
// pre-fix 405-on-cleanup bug, etc.).
//
// Best-effort: any failure here is reported as a warning and does not
// block the run. Single-host scope: the matcher is a description
// prefix, so a future multi-host setup would benefit from a
// host/PID-stamped description to avoid cleaning up other machines'
// active runs.
func cleanupOrphanEndpoints(ctx context.Context, cmd *cli.Command) {
	w := cmd.ErrWriter
	ids, err := listMatchingEndpointIDs(ctx, cmd, defaultDescription)
	if err != nil {
		fmt.Fprintf(w, "warning: could not list existing webhook endpoints (%v); continuing\n", err)
		return
	}
	if len(ids) == 0 {
		return
	}
	deleted := 0
	for _, id := range ids {
		if err := deleteEndpoint(ctx, cmd, id); err != nil {
			fmt.Fprintf(w, "warning: failed to delete stale endpoint %s: %v\n", id, err)
			continue
		}
		deleted++
	}
	if deleted > 0 {
		suffix := ""
		if deleted != 1 {
			suffix = "s"
		}
		fmt.Fprintf(w, "✓ Cleaned up %d stale webhook endpoint%s from prior runs\n", deleted, suffix)
	}
}

func listMatchingEndpointIDs(ctx context.Context, cmd *cli.Command, descriptionMatch string) ([]string, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	baseURL := strings.TrimRight(cmd.String("base-url"), "/")
	if baseURL == "" {
		baseURL = "https://api.case.dev"
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/webhooks/v1/endpoints", nil)
	if err != nil {
		return nil, err
	}
	if apiKey := cmd.String("api-key"); apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read list response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return parseListEndpointsResponse(body, descriptionMatch)
}

// parseListEndpointsResponse extracts ids from a list-endpoints
// response, restricted to entries whose description matches
// descriptionMatch. It tolerates the same wrapper variations as
// parseCreateEndpointResponse: the array can be top-level or nested
// under "endpoints", "data", "results", or "items".
func parseListEndpointsResponse(body []byte, descriptionMatch string) ([]string, error) {
	var raw any
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("decode list response: %w (body: %s)", err, strings.TrimSpace(string(body)))
	}
	items := findEndpointsArray(raw)
	if items == nil {
		return nil, nil
	}
	var ids []string
	for _, item := range items {
		entry, ok := item.(map[string]any)
		if !ok {
			continue
		}
		// The list response may put endpoint fields directly on the
		// entry, or nested under "endpoint" (mirroring the create shape).
		sources := []map[string]any{entry}
		if nested, ok := entry["endpoint"].(map[string]any); ok {
			sources = append(sources, nested)
		}
		desc := ""
		for _, src := range sources {
			if v, ok := src["description"].(string); ok {
				desc = v
				break
			}
		}
		if desc != descriptionMatch {
			continue
		}
		for _, src := range sources {
			if id, ok := src["id"].(string); ok && id != "" {
				ids = append(ids, id)
				break
			}
			if id, ok := src["endpointId"].(string); ok && id != "" {
				ids = append(ids, id)
				break
			}
		}
	}
	return ids, nil
}

func findEndpointsArray(raw any) []any {
	if arr, ok := raw.([]any); ok {
		return arr
	}
	m, ok := raw.(map[string]any)
	if !ok {
		return nil
	}
	for _, key := range []string{"endpoints", "data", "results", "items"} {
		if arr, ok := m[key].([]any); ok {
			return arr
		}
	}
	return nil
}
