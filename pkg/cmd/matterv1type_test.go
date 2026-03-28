// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1TypesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:types", "create",
			"--name", "name",
			"--default-agent-type-id", "default_agent_type_id",
			"--default-metadata", "{foo: bar}",
			"--default-work-item", "{foo: bar}",
			"--description", "description",
			"--exit-criterion", "string",
			"--instructions", "instructions",
			"--intake-requirement", "string",
			"--is-active=true",
			"--orchestration-mode", "auto",
			"--review-agent-type-id", "review_agent_type_id",
			"--review-criterion", "string",
			"--skill-ref", "string",
			"--slug", "slug",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"default_agent_type_id: default_agent_type_id\n" +
			"default_metadata:\n" +
			"  foo: bar\n" +
			"default_work_items:\n" +
			"  - foo: bar\n" +
			"description: description\n" +
			"exit_criteria:\n" +
			"  - string\n" +
			"instructions: instructions\n" +
			"intake_requirements:\n" +
			"  - string\n" +
			"is_active: true\n" +
			"orchestration_mode: auto\n" +
			"review_agent_type_id: review_agent_type_id\n" +
			"review_criteria:\n" +
			"  - string\n" +
			"skill_refs:\n" +
			"  - string\n" +
			"slug: slug\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:types", "create",
		)
	})
}

func TestMattersV1TypesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:types", "retrieve",
			"--id", "id",
		)
	})
}

func TestMattersV1TypesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:types", "update",
			"--id", "id",
			"--default-agent-type-id", "default_agent_type_id",
			"--default-metadata", "{foo: bar}",
			"--default-work-item", "{foo: bar}",
			"--description", "description",
			"--exit-criterion", "string",
			"--instructions", "instructions",
			"--intake-requirement", "string",
			"--is-active=true",
			"--name", "name",
			"--orchestration-mode", "auto",
			"--review-agent-type-id", "review_agent_type_id",
			"--review-criterion", "string",
			"--skill-ref", "string",
			"--slug", "slug",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"default_agent_type_id: default_agent_type_id\n" +
			"default_metadata:\n" +
			"  foo: bar\n" +
			"default_work_items:\n" +
			"  - foo: bar\n" +
			"description: description\n" +
			"exit_criteria:\n" +
			"  - string\n" +
			"instructions: instructions\n" +
			"intake_requirements:\n" +
			"  - string\n" +
			"is_active: true\n" +
			"name: name\n" +
			"orchestration_mode: auto\n" +
			"review_agent_type_id: review_agent_type_id\n" +
			"review_criteria:\n" +
			"  - string\n" +
			"skill_refs:\n" +
			"  - string\n" +
			"slug: slug\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:types", "update",
			"--id", "id",
		)
	})
}

func TestMattersV1TypesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:types", "list",
			"--active=true",
		)
	})
}
