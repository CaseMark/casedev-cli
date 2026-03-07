// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentV1ExecuteCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:execute", "create",
		"--api-key", "string",
		"--prompt", "prompt",
		"--disabled-tool", "[string]",
		"--enabled-tool", "[string]",
		"--guidance", "guidance",
		"--instructions", "instructions",
		"--model", "model",
		"--object-id", "[string]",
		"--sandbox", "{cpu: 0, memoryMiB: 0}",
		"--vault-id", "[string]",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(agentV1ExecuteCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"agent:v1:execute", "create",
		"--api-key", "string",
		"--prompt", "prompt",
		"--disabled-tool", "[string]",
		"--enabled-tool", "[string]",
		"--guidance", "guidance",
		"--instructions", "instructions",
		"--model", "model",
		"--object-id", "[string]",
		"--sandbox.cpu", "0",
		"--sandbox.memory-mi-b", "0",
		"--vault-id", "[string]",
	)
}
