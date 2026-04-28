package webhooklisten

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/urfave/cli/v3"
)

const configRelPath = ".casedev/webhooks.json"

var errConfigNotFound = errors.New("webhook listen config not found")

type configFile struct {
	WebhookListen config `json:"webhookListen"`
}

type config struct {
	Events    []string `json:"events,omitempty"`
	ForwardTo string   `json:"forwardTo,omitempty"`
}

type resolvedConfig struct {
	Events        []string
	ForwardTo     string
	ProjectPath   string
	GlobalPath    string
	ProjectExists bool
	GlobalExists  bool
	EventsSources map[string]string
	ForwardSource string
}

func projectConfigPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, configRelPath), nil
}

func globalConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "casedev", "webhooks.json"), nil
}

func loadConfigFromPath(path string) (config, bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return config{}, false, nil
		}
		return config{}, false, err
	}

	var cfg configFile
	if err := json.Unmarshal(data, &cfg); err != nil {
		return config{}, true, fmt.Errorf("parse %s: %w", path, err)
	}

	cfg.WebhookListen.Events = normalizePatterns(cfg.WebhookListen.Events)
	cfg.WebhookListen.ForwardTo = strings.TrimSpace(cfg.WebhookListen.ForwardTo)
	return cfg.WebhookListen, true, nil
}

func loadConfig(globalOnly bool) (resolvedConfig, error) {
	projectPath, err := projectConfigPath()
	if err != nil {
		return resolvedConfig{}, err
	}
	globalPath, err := globalConfigPath()
	if err != nil {
		return resolvedConfig{}, err
	}

	projectCfg, projectExists, err := loadConfigFromPath(projectPath)
	if err != nil {
		return resolvedConfig{}, err
	}
	globalCfg, globalExists, err := loadConfigFromPath(globalPath)
	if err != nil {
		return resolvedConfig{}, err
	}

	resolved := resolvedConfig{
		ProjectPath:   projectPath,
		GlobalPath:    globalPath,
		ProjectExists: projectExists,
		GlobalExists:  globalExists,
		EventsSources: map[string]string{},
	}

	if globalOnly {
		resolved.Events = append([]string(nil), globalCfg.Events...)
		for _, pattern := range resolved.Events {
			resolved.EventsSources[pattern] = fmt.Sprintf("global: %s", globalPath)
		}
		resolved.ForwardTo = globalCfg.ForwardTo
		if globalCfg.ForwardTo != "" {
			resolved.ForwardSource = fmt.Sprintf("global: %s", globalPath)
		}
		if !globalExists {
			return resolved, errConfigNotFound
		}
		return resolved, nil
	}

	resolved.Events = append([]string(nil), globalCfg.Events...)
	for _, pattern := range resolved.Events {
		resolved.EventsSources[pattern] = fmt.Sprintf("global: %s", globalPath)
	}
	resolved.ForwardTo = globalCfg.ForwardTo
	if globalCfg.ForwardTo != "" {
		resolved.ForwardSource = fmt.Sprintf("global: %s", globalPath)
	}

	for _, pattern := range projectCfg.Events {
		if !slices.Contains(resolved.Events, pattern) {
			resolved.Events = append(resolved.Events, pattern)
		}
		resolved.EventsSources[pattern] = fmt.Sprintf("project: %s", projectPath)
	}
	if projectCfg.ForwardTo != "" {
		resolved.ForwardTo = projectCfg.ForwardTo
		resolved.ForwardSource = fmt.Sprintf("project: %s", projectPath)
	}

	if !projectExists && !globalExists {
		return resolved, errConfigNotFound
	}

	return resolved, nil
}

func saveConfig(path string, cfg config) error {
	cfg.Events = normalizePatterns(cfg.Events)
	cfg.ForwardTo = strings.TrimSpace(cfg.ForwardTo)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	payload := configFile{WebhookListen: cfg}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

func normalizePatterns(patterns []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if _, ok := seen[pattern]; ok {
			continue
		}
		seen[pattern] = struct{}{}
		out = append(out, pattern)
	}
	return out
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

func handleInit(_ context.Context, cmd *cli.Command) error {
	path, err := projectConfigPath()
	if err != nil {
		return err
	}
	if cmd.Bool("global") {
		path, err = globalConfigPath()
		if err != nil {
			return err
		}
	}

	cfg, _, err := loadConfigFromPath(path)
	if err != nil {
		return err
	}
	cfg.Events = normalizePatterns(cfg.Events)

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

	if err := saveConfig(path, cfg); err != nil {
		return err
	}

	fmt.Fprintf(cmd.Writer, "Saved webhook listen config to %s\n", path)
	return nil
}

func handleShow(_ context.Context, cmd *cli.Command) error {
	cfg, err := loadConfig(cmd.Bool("global"))
	if err != nil {
		if errors.Is(err, errConfigNotFound) {
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
