package cmd

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
	"slices"
	"strings"
	"syscall"
	"time"

	"github.com/urfave/cli/v3"
)

const webhookListenDefaultDescription = "casedev webhook listen (ephemeral ngrok tunnel)"

type webhookEndpointCreateRequest struct {
	URL              string   `json:"url"`
	Description      string   `json:"description,omitempty"`
	EventTypeFilters []string `json:"eventTypeFilters"`
}

type webhookEndpointCreateResponse struct {
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

type webhookListenRuntime struct {
	printMode string
	forwardTo string
	stderr    io.Writer
	http      *http.Client
}

var webhookListenRun = cli.Command{
	Name:            "run",
	Usage:           "Start an ngrok-backed webhook listener and optionally forward deliveries locally.",
	Suggest:         true,
	HideHelpCommand: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "events",
			Usage: "Comma-separated event-type glob patterns",
		},
		&cli.StringFlag{
			Name:      "forward-to",
			Usage:     "Forward received webhooks to this local URL",
			Validator: func(value string) error { return validateOptionalWebhookListenURL(value, "--forward-to") },
		},
		&cli.IntFlag{
			Name:  "tunnel-port",
			Usage: "Local port for the ngrok tunnel (default: auto-assign)",
		},
		&cli.StringFlag{
			Name:  "print",
			Usage: "Output format for received webhooks (pretty, json, silent)",
			Value: "pretty",
			Validator: func(value string) error {
				switch strings.ToLower(strings.TrimSpace(value)) {
				case "pretty", "json", "silent":
					return nil
				default:
					return fmt.Errorf("--print must be one of: pretty, json, silent")
				}
			},
		},
	},
	Action: handleWebhookListenRun,
}

var webhookListenInit = cli.Command{
	Name:            "init",
	Usage:           "Create or update saved webhook listen configuration.",
	Suggest:         true,
	HideHelpCommand: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "global", Aliases: []string{"g"}, Usage: "Use ~/.config/casedev/webhooks.json instead of .casedev/webhooks.json"},
		&cli.StringSliceFlag{Name: "add", Usage: "Add an event pattern to saved config"},
		&cli.StringSliceFlag{Name: "remove", Usage: "Remove an event pattern from saved config"},
		&cli.StringFlag{
			Name:      "endpoint",
			Usage:     "Set the default forward-to URL",
			Validator: func(value string) error { return validateOptionalWebhookListenURL(value, "--endpoint") },
		},
	},
	Action: handleWebhookListenInit,
}

var webhookListenShow = cli.Command{
	Name:            "show",
	Usage:           "Print resolved webhook listen configuration.",
	Suggest:         true,
	HideHelpCommand: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "global", Aliases: []string{"g"}, Usage: "Show only ~/.config/casedev/webhooks.json"},
	},
	Action: handleWebhookListenShow,
}

func handleWebhookListenRun(ctx context.Context, cmd *cli.Command) error {
	if _, err := exec.LookPath("ngrok"); err != nil {
		return fmt.Errorf("ngrok not found in PATH. Install it first (for example: brew install ngrok)")
	}

	cfg, err := loadWebhookListenConfig(false)
	if err != nil && !errors.Is(err, errWebhookListenConfigNotFound) {
		return err
	}

	events := cfg.Events
	if raw := strings.TrimSpace(cmd.String("events")); raw != "" {
		events = normalizeWebhookListenPatterns(strings.Split(raw, ","))
	}
	if len(events) == 0 {
		events = []string{"*"}
	}

	forwardTo := cfg.ForwardTo
	if cmd.IsSet("forward-to") {
		forwardTo = strings.TrimSpace(cmd.String("forward-to"))
	}
	printMode := strings.ToLower(strings.TrimSpace(cmd.String("print")))

	listener, port, err := startWebhookListenHTTPServer(cmd.Int("tunnel-port"), printMode, forwardTo)
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

	endpoint, err := createWebhookListenEndpoint(ctx, cmd, tunnelURL, events)
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
	if err := deleteWebhookListenEndpoint(cleanupCtx, cmd, endpoint.ID); err != nil {
		fmt.Fprintf(cmd.ErrWriter, "warning: failed to delete webhook endpoint %s: %v\n", endpoint.ID, err)
	} else {
		fmt.Fprintf(cmd.ErrWriter, "✓ Deleted webhook endpoint %s\n", endpoint.ID)
	}
	fmt.Fprintln(cmd.ErrWriter, "✓ Tunnel closed")
	return nil
}

func handleWebhookListenInit(ctx context.Context, cmd *cli.Command) error {
	_ = ctx
	path, err := webhookListenProjectConfigPath()
	if err != nil {
		return err
	}
	if cmd.Bool("global") {
		path, err = webhookListenGlobalConfigPath()
		if err != nil {
			return err
		}
	}

	cfg, _, err := loadWebhookListenConfigFromPath(path)
	if err != nil {
		return err
	}
	cfg.Events = normalizeWebhookListenPatterns(cfg.Events)

	for _, pattern := range cmd.StringSlice("add") {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if !slices.Contains(cfg.Events, pattern) {
			cfg.Events = append(cfg.Events, pattern)
		}
	}

	for _, pattern := range cmd.StringSlice("remove") {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		cfg.Events = deletePattern(cfg.Events, pattern)
	}

	if cmd.IsSet("endpoint") {
		cfg.ForwardTo = strings.TrimSpace(cmd.String("endpoint"))
	}

	if err := saveWebhookListenConfig(path, cfg); err != nil {
		return err
	}

	fmt.Fprintf(cmd.Writer, "Saved webhook listen config to %s\n", path)
	return nil
}

func handleWebhookListenShow(ctx context.Context, cmd *cli.Command) error {
	_ = ctx
	cfg, err := loadWebhookListenConfig(cmd.Bool("global"))
	if err != nil {
		if errors.Is(err, errWebhookListenConfigNotFound) {
			fmt.Fprintln(cmd.Writer, "No webhook listen configuration found.")
			return nil
		}
		return err
	}

	fmt.Fprintln(cmd.Writer, "Webhook Listen Configuration")
	fmt.Fprintln(cmd.Writer)
	fmt.Fprintln(cmd.Writer, "  Events:")
	if len(cfg.Events) == 0 {
		fmt.Fprintln(cmd.Writer, "    (none)")
	} else {
		for _, pattern := range cfg.Events {
			source := cfg.EventsSources[pattern]
			if source == "" {
				source = "runtime default"
			}
			fmt.Fprintf(cmd.Writer, "    %s    (%s)\n", pattern, source)
		}
	}
	fmt.Fprintln(cmd.Writer)
	if cfg.ForwardTo == "" {
		fmt.Fprintln(cmd.Writer, "  Forward to: (not set)")
	} else {
		fmt.Fprintf(cmd.Writer, "  Forward to: %s  (%s)\n", cfg.ForwardTo, cfg.ForwardSource)
	}
	fmt.Fprintln(cmd.Writer)
	fmt.Fprintln(cmd.Writer, "  Config files:")
	fmt.Fprintf(cmd.Writer, "    Project: %s %s\n", cfg.ProjectPath, existsMarker(cfg.ProjectExists))
	fmt.Fprintf(cmd.Writer, "    Global:  %s %s\n", cfg.GlobalPath, existsMarker(cfg.GlobalExists))
	return nil
}

func validateOptionalWebhookListenURL(value string, flagName string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return ValidateBaseURL(value, flagName)
}

func deletePattern(patterns []string, target string) []string {
	filtered := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		if pattern != target {
			filtered = append(filtered, pattern)
		}
	}
	return filtered
}

func existsMarker(ok bool) string {
	if ok {
		return "✓"
	}
	return "-"
}

func startWebhookListenHTTPServer(port int, printMode string, forwardTo string) (*webhookListenServer, int, error) {
	addr := "127.0.0.1:0"
	if port > 0 {
		addr = fmt.Sprintf("127.0.0.1:%d", port)
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, 0, err
	}
	port := ln.Addr().(*net.TCPAddr).Port
	server := &webhookListenServer{
		listener: ln,
		runtime: webhookListenRuntime{
			printMode: printMode,
			forwardTo: forwardTo,
			stderr:    os.Stderr,
			http:      &http.Client{Timeout: 15 * time.Second},
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.handleWebhook)
	server.server = &http.Server{Handler: mux}
	go func() {
		_ = server.server.Serve(ln)
	}()
	return server, port, nil
}

type webhookListenServer struct {
	listener net.Listener
	server   *http.Server
	runtime  webhookListenRuntime
}

func (s *webhookListenServer) close() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = s.server.Shutdown(ctx)
}

func (s *webhookListenServer) handleWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	_ = r.Body.Close()

	printWebhookListenEvent(s.runtime, body)

	statusLine := ""
	if s.runtime.forwardTo != "" {
		statusLine = forwardWebhookListenRequest(s.runtime, r, body)
	}
	if statusLine != "" && s.runtime.printMode == "pretty" {
		fmt.Fprintln(s.runtime.stderr, statusLine)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func printWebhookListenEvent(runtime webhookListenRuntime, body []byte) {
	if runtime.printMode == "silent" {
		return
	}
	if runtime.printMode == "json" {
		fmt.Fprintln(runtime.stderr, string(body))
		return
	}

	var payload struct {
		Type       string `json:"type"`
		OccurredAt string `json:"occurred_at"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Fprintf(runtime.stderr, "%s\n", string(body))
		return
	}
	if payload.Type == "" {
		fmt.Fprintf(runtime.stderr, "%s\n", string(body))
		return
	}
	if payload.OccurredAt != "" {
		fmt.Fprintf(runtime.stderr, "%s → %s\n", payload.OccurredAt, payload.Type)
		return
	}
	fmt.Fprintf(runtime.stderr, "%s\n", payload.Type)
}

func forwardWebhookListenRequest(runtime webhookListenRuntime, incoming *http.Request, body []byte) string {
	req, err := http.NewRequestWithContext(incoming.Context(), incoming.Method, runtime.forwardTo, bytes.NewReader(body))
	if err != nil {
		return fmt.Sprintf("warning: forward failed: %v", err)
	}
	for key, values := range incoming.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	resp, err := runtime.http.Do(req)
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

func createWebhookListenEndpoint(ctx context.Context, cmd *cli.Command, tunnelURL string, events []string) (webhookEndpointCreateResponse, error) {
	payload := webhookEndpointCreateRequest{
		URL:              tunnelURL,
		Description:      webhookListenDefaultDescription,
		EventTypeFilters: events,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return webhookEndpointCreateResponse{}, err
	}
	client := &http.Client{Timeout: 15 * time.Second}
	baseURL := strings.TrimRight(cmd.String("base-url"), "/")
	if baseURL == "" {
		baseURL = "https://api.case.dev"
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/webhooks/v1/endpoints", bytes.NewReader(body))
	if err != nil {
		return webhookEndpointCreateResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if apiKey := cmd.String("api-key"); apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return webhookEndpointCreateResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		return webhookEndpointCreateResponse{}, fmt.Errorf("create webhook endpoint: %s: %s", resp.Status, strings.TrimSpace(string(data)))
	}
	var endpoint webhookEndpointCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&endpoint); err != nil {
		return webhookEndpointCreateResponse{}, err
	}
	return endpoint, nil
}

func deleteWebhookListenEndpoint(ctx context.Context, cmd *cli.Command, endpointID string) error {
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
