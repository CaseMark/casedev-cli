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
