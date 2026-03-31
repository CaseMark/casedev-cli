// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentV2ExecuteCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:execute", "create",
			"--prompt", "prompt",
			"--agent-runtime=true",
			"--disabled-tool", "[string]",
			"--enabled-tool", "[string]",
			"--guidance", "guidance",
			"--instructions", "instructions",
			"--model", "model",
			"--object-id", "[string]",
			"--sandbox", "{cpu: 0, memoryMiB: 0}",
			"--vault-id", "[string]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentV2ExecuteCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:execute", "create",
			"--prompt", "prompt",
			"--agent-runtime=true",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"prompt: prompt\n" +
			"agentRuntime: true\n" +
			"disabledTools:\n" +
			"  - string\n" +
			"enabledTools:\n" +
			"  - string\n" +
			"guidance: guidance\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"objectIds:\n" +
			"  - string\n" +
			"sandbox:\n" +
			"  cpu: 0\n" +
			"  memoryMiB: 0\n" +
			"vaultIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:v2:execute", "create",
		)
	})
}
