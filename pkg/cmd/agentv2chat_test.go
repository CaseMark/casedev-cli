// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentV2ChatCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "create",
			"--idle-timeout-ms", "0",
			"--model", "model",
			"--title", "title",
			"--vault-id", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"idleTimeoutMs: 0\n" +
			"model: model\n" +
			"title: title\n" +
			"vaultIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v2:chat", "create",
		)
	})
}

func TestAgentV2ChatDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "delete",
			"--id", "id",
		)
	})
}

func TestAgentV2ChatCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "cancel",
			"--id", "id",
		)
	})
}

func TestAgentV2ChatReplyToQuestion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "reply-to-question",
			"--id", "id",
			"--request-id", "requestID",
			"--answer", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"answers:\n" +
			"  - - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v2:chat", "reply-to-question",
			"--id", "id",
			"--request-id", "requestID",
		)
	})
}

func TestAgentV2ChatRespond(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "respond",
			"--max-items", "10",
			"--id", "id",
			"--part", "{text: text, type: text}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentV2ChatRespond)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "respond",
			"--max-items", "10",
			"--id", "id",
			"--part.text", "text",
			"--part.type", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"parts:\n" +
			"  - text: text\n" +
			"    type: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v2:chat", "respond",
			"--max-items", "10",
			"--id", "id",
		)
	})
}

func TestAgentV2ChatSendMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "send-message",
			"--id", "id",
			"--part", "{text: text, type: text}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentV2ChatSendMessage)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "send-message",
			"--id", "id",
			"--part.text", "text",
			"--part.type", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"parts:\n" +
			"  - text: text\n" +
			"    type: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v2:chat", "send-message",
			"--id", "id",
		)
	})
}

func TestAgentV2ChatStream(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat", "stream",
			"--max-items", "10",
			"--id", "id",
			"--last-event-id", "0",
		)
	})
}
