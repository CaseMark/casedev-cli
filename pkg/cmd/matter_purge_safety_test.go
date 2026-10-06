package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatterPurgeSafety(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "casedev")
	build := exec.Command("go", "build", "-o", binary, "../../cmd/casedev")
	output, err := build.CombinedOutput()
	require.NoError(t, err, "%s", output)

	for _, test := range []struct {
		name  string
		args  []string
		stdin string
		error string
	}{
		{name: "matter requires confirmation", args: []string{"matters:v1", "delete", "--id", "matter"}, error: "requires --confirm"},
		{name: "false confirmation is refused", args: []string{"matters:v1", "delete", "--id", "matter", "--confirm=false"}, error: "requires --confirm"},
		{name: "confirmed matter deletion", args: []string{"matters:v1", "delete", "--id", "matter", "--confirm"}},
		{name: "content requires confirmation", args: []string{"matters:v1:content-purges", "create", "--id", "matter", "--request-id", "request", "--object-id", "object"}, error: "requires --confirm"},
		{name: "empty targets are refused", args: []string{"matters:v1:content-purges", "create", "--id", "matter", "--request-id", "request", "--confirm"}, error: "requires at least one nonempty"},
		{name: "empty stdin targets are refused", args: []string{"matters:v1:content-purges", "create", "--id", "matter", "--confirm"}, stdin: `{"request_id":"request","object_ids":[],"session_ids":[""]}`, error: "requires at least one nonempty"},
		{name: "scalar stdin targets are refused", args: []string{"matters:v1:content-purges", "create", "--id", "matter", "--confirm"}, stdin: `{"request_id":"request","object_ids":"object"}`, error: "requires at least one nonempty"},
		{name: "confirmed flag target", args: []string{"matters:v1:content-purges", "create", "--id", "matter", "--request-id", "request", "--object-id", "object", "--confirm"}},
		{name: "confirmed stdin target", args: []string{"matters:v1:content-purges", "create", "--id", "matter", "--confirm"}, stdin: `{"request_id":"request","object_ids":["object"]}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			var requests atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				if r.Method == http.MethodPost {
					var body map[string]any
					if !assert.NoError(t, json.NewDecoder(r.Body).Decode(&body)) {
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					assert.NotContains(t, body, "confirm")
					assert.Equal(t, "request", body["request_id"])
					assert.Equal(t, []any{"object"}, body["object_ids"])
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{}`))
			}))
			defer server.Close()

			args := append([]string{"--base-url", server.URL, "--api-key", "test-key"}, test.args...)
			command := exec.Command(binary, args...)
			command.Stdin = strings.NewReader(test.stdin)
			output, err := command.CombinedOutput()
			if test.error != "" {
				require.Error(t, err)
				require.Contains(t, string(output), test.error)
				require.Zero(t, requests.Load(), "refused purges must not send requests")
			} else {
				require.NoError(t, err, "%s", output)
				require.EqualValues(t, 1, requests.Load())
			}
		})
	}
}
