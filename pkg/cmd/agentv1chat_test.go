// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV1ChatCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:chat", "create",
		"--idle-timeout-ms", "0",
		"--model", "model",
		"--title", "title",
	)
}

func TestAgentV1ChatDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:chat", "delete",
		"--id", "id",
	)
}

func TestAgentV1ChatCancel(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:chat", "cancel",
		"--id", "id",
	)
}

func TestAgentV1ChatSendMessage(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:chat", "send-message",
		"--id", "id",
		"--body", "{}",
	)
}

func TestAgentV1ChatStream(t *testing.T) {
	t.Skip("Mock server doesn't support text/event-stream responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:chat", "stream",
		"--id", "id",
		"--last-event-id", "0",
	)
}
