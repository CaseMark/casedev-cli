// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV2RunCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:run", "create",
			"--agent-id", "agentId",
			"--prompt", "prompt",
			"--callback-url", "https://example.com",
			"--guidance", "guidance",
			"--model", "model",
			"--object-id", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agentId: agentId\n" +
			"prompt: prompt\n" +
			"callbackUrl: https://example.com\n" +
			"guidance: guidance\n" +
			"model: model\n" +
			"objectIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v2:run", "create",
		)
	})
}

func TestAgentV2RunEvents(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:run", "events",
			"--max-items", "10",
			"--id", "id",
			"--last-event-id", "0",
		)
	})
}

func TestAgentV2RunExec(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:run", "exec",
			"--id", "id",
		)
	})
}

func TestAgentV2RunGetDetails(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:run", "get-details",
			"--id", "id",
		)
	})
}

func TestAgentV2RunGetStatus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:run", "get-status",
			"--id", "id",
		)
	})
}
