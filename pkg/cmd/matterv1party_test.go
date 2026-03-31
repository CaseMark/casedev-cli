// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1PartiesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:parties", "create",
			"--name", "name",
			"--address", "{foo: bar}",
			"--custom-fields", "{foo: bar}",
			"--email", "email",
			"--metadata", "{foo: bar}",
			"--notes", "notes",
			"--phone", "phone",
			"--type", "person",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"addresses:\n" +
			"  - foo: bar\n" +
			"custom_fields:\n" +
			"  foo: bar\n" +
			"email: email\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"notes: notes\n" +
			"phone: phone\n" +
			"type: person\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:parties", "create",
		)
	})
}

func TestMattersV1PartiesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:parties", "retrieve",
			"--party-id", "partyId",
		)
	})
}

func TestMattersV1PartiesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:parties", "update",
			"--party-id", "partyId",
		)
	})
}

func TestMattersV1PartiesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:parties", "list",
			"--email", "email",
			"--query", "query",
			"--type", "person",
		)
	})
}
