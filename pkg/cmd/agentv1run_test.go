// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV1RunCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "create",
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
			"agent:v1:run", "create",
		)
	})
}

func TestAgentV1RunList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "list",
			"--agent-id", "agentId",
			"--cursor", "cursor",
			"--limit", "1",
			"--status", "queued",
		)
	})
}

func TestAgentV1RunCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "cancel",
			"--id", "id",
		)
	})
}

func TestAgentV1RunEvents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "events",
			"--max-items", "10",
			"--id", "id",
			"--last-event-id", "0",
		)
	})
}

func TestAgentV1RunExec(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "exec",
			"--id", "id",
		)
	})
}

func TestAgentV1RunGetDetails(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "get-details",
			"--id", "id",
		)
	})
}

func TestAgentV1RunGetStatus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "get-status",
			"--id", "id",
		)
	})
}

func TestAgentV1RunWatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v1:run", "watch",
			"--id", "id",
			"--callback-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("callbackUrl: https://example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v1:run", "watch",
			"--id", "id",
		)
	})
}
