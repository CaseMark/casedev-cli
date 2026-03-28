// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1MatterPartiesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:matter-parties", "create",
			"--id", "id",
			"--party-id", "party_id",
			"--role", "client",
			"--custom-fields", "{foo: bar}",
			"--is-primary=true",
			"--metadata", "{foo: bar}",
			"--notes", "notes",
			"--set-as-client=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"party_id: party_id\n" +
			"role: client\n" +
			"custom_fields:\n" +
			"  foo: bar\n" +
			"is_primary: true\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"notes: notes\n" +
			"set_as_client: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:matter-parties", "create",
			"--id", "id",
		)
	})
}

func TestMattersV1MatterPartiesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:matter-parties", "list",
			"--id", "id",
		)
	})
}
