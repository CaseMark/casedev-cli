// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV1RunCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:run", "create",
		"--agent-id", "agentId",
		"--prompt", "prompt",
		"--guidance", "guidance",
		"--model", "model",
	)
}

func TestAgentV1RunCancel(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:run", "cancel",
		"--id", "id",
	)
}

func TestAgentV1RunExec(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:run", "exec",
		"--id", "id",
	)
}

func TestAgentV1RunGetDetails(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:run", "get-details",
		"--id", "id",
	)
}

func TestAgentV1RunGetStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:run", "get-status",
		"--id", "id",
	)
}

func TestAgentV1RunWatch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:run", "watch",
		"--id", "id",
		"--callback-url", "https://example.com",
	)
}
