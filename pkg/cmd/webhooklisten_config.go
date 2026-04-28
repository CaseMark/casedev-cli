package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const webhookListenConfigRelPath = ".casedev/webhooks.json"

var errWebhookListenConfigNotFound = errors.New("webhook listen config not found")

type webhookListenConfigFile struct {
	WebhookListen webhookListenConfig `json:"webhookListen"`
}

type webhookListenConfig struct {
	Events    []string `json:"events,omitempty"`
	ForwardTo string   `json:"forwardTo,omitempty"`
}

type resolvedWebhookListenConfig struct {
	Events         []string
	ForwardTo      string
	ProjectPath    string
	GlobalPath     string
	ProjectExists  bool
	GlobalExists   bool
	EventsSources  map[string]string
	ForwardSource  string
}

func webhookListenProjectConfigPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, webhookListenConfigRelPath), nil
}

func webhookListenGlobalConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "casedev", "webhooks.json"), nil
}

func loadWebhookListenConfigFromPath(path string) (webhookListenConfig, bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return webhookListenConfig{}, false, nil
		}
		return webhookListenConfig{}, false, err
	}

	var cfg webhookListenConfigFile
	if err := json.Unmarshal(data, &cfg); err != nil {
		return webhookListenConfig{}, true, fmt.Errorf("parse %s: %w", path, err)
	}

	cfg.WebhookListen.Events = normalizeWebhookListenPatterns(cfg.WebhookListen.Events)
	cfg.WebhookListen.ForwardTo = strings.TrimSpace(cfg.WebhookListen.ForwardTo)
	return cfg.WebhookListen, true, nil
}

func loadWebhookListenConfig(globalOnly bool) (resolvedWebhookListenConfig, error) {
	projectPath, err := webhookListenProjectConfigPath()
	if err != nil {
		return resolvedWebhookListenConfig{}, err
	}
	globalPath, err := webhookListenGlobalConfigPath()
	if err != nil {
		return resolvedWebhookListenConfig{}, err
	}

	projectCfg, projectExists, err := loadWebhookListenConfigFromPath(projectPath)
	if err != nil {
		return resolvedWebhookListenConfig{}, err
	}
	globalCfg, globalExists, err := loadWebhookListenConfigFromPath(globalPath)
	if err != nil {
		return resolvedWebhookListenConfig{}, err
	}

	resolved := resolvedWebhookListenConfig{
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
			return resolved, errWebhookListenConfigNotFound
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
		return resolved, errWebhookListenConfigNotFound
	}

	return resolved, nil
}

func saveWebhookListenConfig(path string, cfg webhookListenConfig) error {
	cfg.Events = normalizeWebhookListenPatterns(cfg.Events)
	cfg.ForwardTo = strings.TrimSpace(cfg.ForwardTo)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	payload := webhookListenConfigFile{WebhookListen: cfg}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

func normalizeWebhookListenPatterns(patterns []string) []string {
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
