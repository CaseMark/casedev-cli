// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentV1ChatCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "create",
			"--api-key", "string",
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
			t, pipeData, "agent:v1:chat", "create",
			"--api-key", "string",
		)
	})
}

func TestAgentV1ChatDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1ChatCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "cancel",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1ChatReplyToQuestion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "reply-to-question",
			"--api-key", "string",
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
			t, pipeData, "agent:v1:chat", "reply-to-question",
			"--api-key", "string",
			"--id", "id",
			"--request-id", "requestID",
		)
	})
}

func TestAgentV1ChatRespond(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "respond",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
			"--part", "{text: text, type: text}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentV1ChatRespond)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "respond",
			"--api-key", "string",
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
			t, pipeData, "agent:v1:chat", "respond",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
		)
	})
}

func TestAgentV1ChatSendMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "send-message",
			"--api-key", "string",
			"--id", "id",
			"--part", "{text: text, type: text}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentV1ChatSendMessage)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "send-message",
			"--api-key", "string",
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
			t, pipeData, "agent:v1:chat", "send-message",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1ChatStream(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat", "stream",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
			"--last-event-id", "0",
		)
	})
}
