// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentV1AgentsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:agents", "create",
		"--instructions", "instructions",
		"--name", "name",
		"--description", "description",
		"--disabled-tool", "[string]",
		"--enabled-tool", "[string]",
		"--model", "model",
		"--sandbox", "{cpu: 0, memoryMiB: 0}",
		"--vault-id", "[string]",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(agentV1AgentsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:agents", "create",
		"--instructions", "instructions",
		"--name", "name",
		"--description", "description",
		"--disabled-tool", "[string]",
		"--enabled-tool", "[string]",
		"--model", "model",
		"--sandbox.cpu", "0",
		"--sandbox.memory-mi-b", "0",
		"--vault-id", "[string]",
	)
}

func TestAgentV1AgentsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:agents", "retrieve",
		"--id", "id",
	)
}

func TestAgentV1AgentsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:agents", "update",
		"--id", "id",
		"--description", "description",
		"--disabled-tool", "[string]",
		"--enabled-tool", "[string]",
		"--instructions", "instructions",
		"--model", "model",
		"--name", "name",
		"--sandbox", "{}",
		"--vault-id", "[string]",
	)
}

func TestAgentV1AgentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:agents", "list",
	)
}

func TestAgentV1AgentsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:agents", "delete",
		"--id", "id",
	)
}
