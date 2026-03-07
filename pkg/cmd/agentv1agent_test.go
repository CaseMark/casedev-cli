// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentV1AgentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:agents", "create",
			"--api-key", "string",
			"--instructions", "instructions",
			"--name", "name",
			"--description", "description",
			"--disabled-tool", "[string]",
			"--enabled-tool", "[string]",
			"--model", "model",
			"--sandbox", "{cpu: 0, memoryMiB: 0}",
			"--vault-group", "[string]",
			"--vault-id", "[string]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentV1AgentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:agents", "create",
			"--api-key", "string",
			"--instructions", "instructions",
			"--name", "name",
			"--description", "description",
			"--disabled-tool", "[string]",
			"--enabled-tool", "[string]",
			"--model", "model",
			"--sandbox.cpu", "0",
			"--sandbox.memory-mi-b", "0",
			"--vault-group", "[string]",
			"--vault-id", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"name: name\n" +
			"description: description\n" +
			"disabledTools:\n" +
			"  - string\n" +
			"enabledTools:\n" +
			"  - string\n" +
			"model: model\n" +
			"sandbox:\n" +
			"  cpu: 0\n" +
			"  memoryMiB: 0\n" +
			"vaultGroups:\n" +
			"  - string\n" +
			"vaultIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "agent:v1:agents", "create",
			"--api-key", "string",
		)
	})
}

func TestAgentV1AgentsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:agents", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1AgentsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:agents", "update",
			"--api-key", "string",
			"--id", "id",
			"--description", "description",
			"--disabled-tool", "[string]",
			"--enabled-tool", "[string]",
			"--instructions", "instructions",
			"--model", "model",
			"--name", "name",
			"--sandbox", "{}",
			"--vault-group", "[string]",
			"--vault-id", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"disabledTools:\n" +
			"  - string\n" +
			"enabledTools:\n" +
			"  - string\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"name: name\n" +
			"sandbox: {}\n" +
			"vaultGroups:\n" +
			"  - string\n" +
			"vaultIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "agent:v1:agents", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1AgentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:agents", "list",
			"--api-key", "string",
		)
	})
}

func TestAgentV1AgentsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:agents", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
