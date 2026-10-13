package webhooklisten

import (
	"bytes"
	"strings"
	"testing"
)

func TestExtractIDs(t *testing.T) {
	cases := []struct {
		name string
		body string
		want string
	}{
		{"top-level id only", `{"id":"evt_1","type":"x","name":"y"}`, "id=evt_1"},
		{"nested data.objectId", `{"type":"x","data":{"objectId":"vobj_1","name":"y"}}`, "data.objectId=vobj_1"},
		{"multiple, sorted", `{"id":"evt_1","data":{"objectId":"vobj_1","ownerId":"u_2"}}`, "data.objectId=vobj_1 data.ownerId=u_2 id=evt_1"},
		{"non-string id ignored", `{"id":42,"data":{"objectId":"vobj_1"}}`, "data.objectId=vobj_1"},
		{"empty payload", `{}`, ""},
		{"invalid json", `not json`, ""},
		{"id-shaped key not at top or known parent", `{"sidebar":{"objectId":"x"}}`, ""},
		{"object parent key recognized", `{"object":{"id":"o_1"}}`, "object.id=o_1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := extractIDs([]byte(tc.body))
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPrintEventPretty(t *testing.T) {
	var buf bytes.Buffer
	rt := runtime{printMode: "pretty", stderr: &buf}
	body := []byte(`{"type":"vault.object.created","occurred_at":"2026-04-28T16:00:00Z","data":{"objectId":"vobj_1"}}`)
	printEvent(rt, body)
	got := buf.String()
	for _, want := range []string{"vault.object.created", "2026-04-28T16:00:00Z", "data.objectId=vobj_1"} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q: %q", want, got)
		}
	}
	if strings.Contains(got, "{") {
		t.Errorf("pretty mode should not contain JSON body: %q", got)
	}
}

func TestPrintEventRaw(t *testing.T) {
	var buf bytes.Buffer
	rt := runtime{printMode: "raw", stderr: &buf}
	body := []byte(`{"type":"x","data":{"objectId":"vobj_1"}}`)
	printEvent(rt, body)
	got := buf.String()
	if !strings.Contains(got, `"objectId": "vobj_1"`) {
		t.Errorf("raw mode missing indented JSON: %q", got)
	}
	if !strings.Contains(got, "\n") {
		t.Errorf("raw mode should emit multi-line JSON: %q", got)
	}
}

func TestParseCreateEndpointResponse(t *testing.T) {
	cases := []struct {
		name    string
		body    string
		wantID  string
		wantSec string
		wantErr bool
	}{
		{"top-level id", `{"id":"ep_1","signingSecret":"sec_1"}`, "ep_1", "sec_1", false},
		{"top-level endpointId", `{"endpointId":"ep_2","secret":"sec_2"}`, "ep_2", "sec_2", false},
		{"wrapped in data", `{"data":{"endpointId":"ep_3","signingSecret":"sec_3"}}`, "ep_3", "sec_3", false},
		{"wrapped in endpoint, secret at top-level", `{"endpoint":{"id":"ep_4","description":"x"},"signingSecret":"sec_4"}`, "ep_4", "sec_4", false},
		{"wrapped data with id key", `{"data":{"id":"ep_5"}}`, "ep_5", "", false},
		{"missing id surfaces raw body", `{"foo":"bar"}`, "", "", true},
		{"invalid json surfaces raw body", `not json`, "", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseCreateEndpointResponse([]byte(tc.body))
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tc.wantErr)
			}
			if got.ID != tc.wantID {
				t.Errorf("id = %q, want %q", got.ID, tc.wantID)
			}
			if got.SigningSecret != tc.wantSec {
				t.Errorf("signingSecret = %q, want %q", got.SigningSecret, tc.wantSec)
			}
			if tc.wantErr && err != nil && !strings.Contains(err.Error(), tc.body[:min(len(tc.body), 20)]) {
				t.Errorf("error should include raw body for diagnosis, got: %v", err)
			}
		})
	}
}

func TestParseListEndpointsResponse(t *testing.T) {
	const marker = "casedev webhook listen (ephemeral ngrok tunnel)"
	cases := []struct {
		name string
		body string
		want []string
	}{
		{
			"top-level array, mixed descriptions",
			`[{"id":"a","description":"casedev webhook listen (ephemeral ngrok tunnel)"},{"id":"b","description":"other"}]`,
			[]string{"a"},
		},
		{
			"wrapped in endpoints",
			`{"endpoints":[{"id":"a","description":"casedev webhook listen (ephemeral ngrok tunnel)"}]}`,
			[]string{"a"},
		},
		{
			"wrapped in data",
			`{"data":[{"endpointId":"b","description":"casedev webhook listen (ephemeral ngrok tunnel)"}]}`,
			[]string{"b"},
		},
		{
			"each item itself wrapped in endpoint",
			`{"endpoints":[{"endpoint":{"id":"a","description":"casedev webhook listen (ephemeral ngrok tunnel)"}}]}`,
			[]string{"a"},
		},
		{
			"empty list",
			`{"endpoints":[]}`,
			nil,
		},
		{
			"no marker matches",
			`[{"id":"a","description":"unrelated"}]`,
			nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseListEndpointsResponse([]byte(tc.body), marker)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("ids[%d] = %q, want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}

func TestUseColorRespectsNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	if useColor(&bytes.Buffer{}) {
		t.Error("NO_COLOR should disable color")
	}
}

func TestUseColorRespectsCLIColorForce(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	t.Setenv("CLICOLOR_FORCE", "1")
	if !useColor(&bytes.Buffer{}) {
		t.Error("CLICOLOR_FORCE=1 should force color even on non-TTY writer")
	}
}
