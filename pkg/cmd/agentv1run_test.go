// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV1RunCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "create",
			"--api-key", "string",
			"--agent-id", "agentId",
			"--prompt", "prompt",
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
			"guidance: guidance\n" +
			"model: model\n" +
			"objectIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "agent:v1:run", "create",
			"--api-key", "string",
		)
	})
}

func TestAgentV1RunCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "cancel",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1RunEvents(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "events",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
			"--last-event-id", "0",
		)
	})
}

func TestAgentV1RunExec(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "exec",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1RunGetDetails(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "get-details",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1RunGetStatus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "get-status",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1RunWatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:run", "watch",
			"--api-key", "string",
			"--id", "id",
			"--callback-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("callbackUrl: https://example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "agent:v1:run", "watch",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
