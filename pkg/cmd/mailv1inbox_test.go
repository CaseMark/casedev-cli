// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMailV1InboxesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "create",
			"--api-key", "string",
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
			t, pipeData, "mail:v1:inboxes", "create",
			"--api-key", "string",
		)
	})
}

func TestMailV1InboxesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "retrieve",
			"--api-key", "string",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "list",
			"--api-key", "string",
		)
	})
}

func TestMailV1InboxesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "delete",
			"--api-key", "string",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesGetAttachment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "get-attachment",
			"--api-key", "string",
			"--inbox-id", "inboxId",
			"--message-id", "messageId",
			"--attachment-id", "attachmentId",
		)
	})
}

func TestMailV1InboxesGetMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "get-message",
			"--api-key", "string",
			"--inbox-id", "inboxId",
			"--message-id", "messageId",
		)
	})
}

func TestMailV1InboxesListMessages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "list-messages",
			"--api-key", "string",
			"--inbox-id", "inboxId",
		)
	})
}

func TestMailV1InboxesReply(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "reply",
			"--api-key", "string",
			"--inbox-id", "inboxId",
			"--message-id", "messageId",
		)
	})
}

func TestMailV1InboxesSend(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mail:v1:inboxes", "send",
			"--api-key", "string",
			"--inbox-id", "inboxId",
		)
	})
}
