package webhooklisten

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/charmbracelet/x/term"
	"github.com/tidwall/pretty"
)

const (
	ansiReset   = "\x1b[0m"
	ansiDim     = "\x1b[2m"
	ansiBold    = "\x1b[1m"
	ansiCyan    = "\x1b[36m"
	ansiGreen   = "\x1b[32m"
	ansiMagenta = "\x1b[35m"
)

// printEvent writes a webhook delivery to stderr in the format
// dictated by rt.printMode. Modes:
//
//   - pretty: one-line header (timestamp + type + extracted IDs),
//     ANSI-colored when stderr is a TTY. The default.
//   - raw:    indented JSON of the webhook body, ANSI-colored when
//     stderr is a TTY.
func printEvent(rt runtime, body []byte) {
	color := useColor(rt.stderr)
	switch rt.printMode {
	case "raw":
		fmt.Fprintln(rt.stderr, formatJSON(body, color))
	default:
		fmt.Fprintln(rt.stderr, buildHeader(body, color))
	}
}

// buildHeader returns the "<occurred_at> → <type>  <ids>" line.
func buildHeader(body []byte, color bool) string {
	var meta struct {
		Type       string `json:"type"`
		OccurredAt string `json:"occurred_at"`
	}
	_ = json.Unmarshal(body, &meta)

	var sb strings.Builder
	if meta.OccurredAt != "" {
		writeColor(&sb, ansiDim, meta.OccurredAt, color)
		sb.WriteString(" → ")
	}
	if meta.Type != "" {
		writeColor(&sb, ansiMagenta+ansiBold, meta.Type, color)
	} else if meta.OccurredAt == "" {
		s := strings.TrimSpace(string(body))
		if len(s) > 80 {
			s = s[:80] + "..."
		}
		sb.WriteString(s)
	}
	for i, pair := range extractIDPairs(body) {
		if i == 0 {
			sb.WriteString("  ")
		} else {
			sb.WriteString(" ")
		}
		writeColor(&sb, ansiGreen, pair.key, color)
		sb.WriteString("=")
		writeColor(&sb, ansiCyan, pair.value, color)
	}
	return sb.String()
}

type idPair struct {
	key, value string
}

// extractIDPairs returns the same data extractIDs surfaces but in
// structured form so callers can render keys and values with
// different colors. Pairs are sorted by key.
func extractIDPairs(body []byte) []idPair {
	var raw map[string]any
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil
	}
	found := map[string]string{}
	collect := func(prefix string, m map[string]any) {
		for key, val := range m {
			if !looksLikeIDKey(key) {
				continue
			}
			s, ok := val.(string)
			if !ok || s == "" {
				continue
			}
			label := key
			if prefix != "" {
				label = prefix + "." + key
			}
			found[label] = s
		}
	}
	collect("", raw)
	for _, parent := range []string{"data", "object", "payload", "subject"} {
		if nested, ok := raw[parent].(map[string]any); ok {
			collect(parent, nested)
		}
	}
	if len(found) == 0 {
		return nil
	}
	labels := make([]string, 0, len(found))
	for k := range found {
		labels = append(labels, k)
	}
	sort.Strings(labels)
	pairs := make([]idPair, 0, len(labels))
	for _, label := range labels {
		pairs = append(pairs, idPair{key: label, value: found[label]})
	}
	return pairs
}

// extractIDs is the joined "key=value [key=value]..." string form of
// extractIDPairs. The walker is intentionally heuristic — it surfaces
// any string field whose key is "id" or ends in "Id"/"ID" at the top
// level or inside common data-bearing parents (data, object, payload,
// subject). The casedev SDK does not yet expose typed webhook
// payloads, so we surface whatever ID-shaped fields are present
// rather than encoding event-specific schemas.
func extractIDs(body []byte) string {
	pairs := extractIDPairs(body)
	if len(pairs) == 0 {
		return ""
	}
	parts := make([]string, len(pairs))
	for i, p := range pairs {
		parts[i] = p.key + "=" + p.value
	}
	return strings.Join(parts, " ")
}

func looksLikeIDKey(k string) bool {
	return k == "id" || strings.HasSuffix(k, "Id") || strings.HasSuffix(k, "ID")
}

// formatJSON returns body indented and, on a TTY, ANSI-colored using
// tidwall/pretty's terminal style. Falls back to the raw body if
// pretty-printing fails.
func formatJSON(body []byte, color bool) string {
	indented := pretty.Pretty(body)
	indented = bytes.TrimRight(indented, "\n")
	if len(indented) == 0 {
		return string(body)
	}
	if color {
		indented = pretty.Color(indented, pretty.TerminalStyle)
	}
	return string(indented)
}

// useColor returns true when ANSI color is appropriate for w. Honors
// the NO_COLOR convention (https://no-color.org) and CLICOLOR_FORCE
// for piping to color-aware pagers like `less -R`.
func useColor(w io.Writer) bool {
	if v := os.Getenv("NO_COLOR"); v != "" {
		return false
	}
	if v := os.Getenv("CLICOLOR_FORCE"); v != "" && v != "0" {
		return true
	}
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(f.Fd())
}

func writeColor(sb *strings.Builder, code, value string, enabled bool) {
	if enabled {
		sb.WriteString(code)
	}
	sb.WriteString(value)
	if enabled {
		sb.WriteString(ansiReset)
	}
}
