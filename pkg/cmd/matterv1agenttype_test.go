// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1AgentTypesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:agent-types", "create",
			"--instructions", "instructions",
			"--name", "name",
			"--description", "description",
			"--disabled-tool", "string",
			"--enabled-tool", "string",
			"--is-active=true",
			"--is-default=true",
			"--metadata", "{foo: bar}",
			"--model", "model",
			"--skill-ref", "string",
			"--slug", "slug",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"name: name\n" +
			"description: description\n" +
			"disabled_tools:\n" +
			"  - string\n" +
			"enabled_tools:\n" +
			"  - string\n" +
			"is_active: true\n" +
			"is_default: true\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"model: model\n" +
			"skill_refs:\n" +
			"  - string\n" +
			"slug: slug\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:agent-types", "create",
		)
	})
}

func TestMattersV1AgentTypesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:agent-types", "list",
			"--active=true",
		)
	})
}
