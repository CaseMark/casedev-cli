// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMailV1InboxesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "create",
			"--address", "address",
			"--display-name", "displayName",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address: address\n" +
			"displayName: displayName\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mail:v1:inboxes", "create",
		)
	})
}

func TestMailV1InboxesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "retrieve",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "list",
		)
	})
}

func TestMailV1InboxesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "delete",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesGetAttachment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "get-attachment",
			"--inbox-id", "inboxId",
			"--message-id", "messageId",
			"--attachment-id", "attachmentId",
		)
	})
}

func TestMailV1InboxesGetMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "get-message",
			"--inbox-id", "inboxId",
			"--message-id", "messageId",
		)
	})
}

func TestMailV1InboxesGetPolicy(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "get-policy",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesListMessages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "list-messages",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesReply(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "reply",
			"--inbox-id", "inboxId",
			"--message-id", "messageId",
		)
	})
}

func TestMailV1InboxesSend(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "send",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesSetPolicy(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mail:v1:inboxes", "set-policy",
			"--inbox-id", "inboxId",
			"--allowed-sender-pattern", "string",
			"--enforce-sender-allowlist=true",
			"--read-access-rule", "string",
			"--reply-access-rule", "string",
			"--send-access-rule", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowedSenderPatterns:\n" +
			"  - string\n" +
			"enforceSenderAllowlist: true\n" +
			"readAccessRules:\n" +
			"  - string\n" +
			"replyAccessRules:\n" +
			"  - string\n" +
			"sendAccessRules:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mail:v1:inboxes", "set-policy",
			"--inbox-id", "inboxId",
		)
	})
}
