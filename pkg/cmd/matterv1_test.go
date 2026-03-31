// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestMattersV1Create(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1", "create",
			"--title", "title",
			"--billing", "{foo: bar}",
			"--client-name", "client_name",
			"--client-party-id", "client_party_id",
			"--custom-fields", "{foo: bar}",
			"--description", "description",
			"--display-id", "display_id",
			"--important-dates", "{foo: bar}",
			"--jurisdiction", "{foo: bar}",
			"--matter-type", "matter_type",
			"--metadata", "{foo: bar}",
			"--practice-area", "practice_area",
			"--responsible-attorney-id", "responsible_attorney_id",
			"--status", "intake",
			"--subtype", "subtype",
			"--vault", "{description: description, enableGraph: true, enableIndexing: true, metadata: {foo: bar}}",
			"--vault-id", "vault_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mattersV1Create)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1", "create",
			"--title", "title",
			"--billing", "{foo: bar}",
			"--client-name", "client_name",
			"--client-party-id", "client_party_id",
			"--custom-fields", "{foo: bar}",
			"--description", "description",
			"--display-id", "display_id",
			"--important-dates", "{foo: bar}",
			"--jurisdiction", "{foo: bar}",
			"--matter-type", "matter_type",
			"--metadata", "{foo: bar}",
			"--practice-area", "practice_area",
			"--responsible-attorney-id", "responsible_attorney_id",
			"--status", "intake",
			"--subtype", "subtype",
			"--vault.description", "description",
			"--vault.enable-graph=true",
			"--vault.enable-indexing=true",
			"--vault.metadata", "{foo: bar}",
			"--vault-id", "vault_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"title: title\n" +
			"billing:\n" +
			"  foo: bar\n" +
			"client_name: client_name\n" +
			"client_party_id: client_party_id\n" +
			"custom_fields:\n" +
			"  foo: bar\n" +
			"description: description\n" +
			"display_id: display_id\n" +
			"important_dates:\n" +
			"  foo: bar\n" +
			"jurisdiction:\n" +
			"  foo: bar\n" +
			"matter_type: matter_type\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"practice_area: practice_area\n" +
			"responsible_attorney_id: responsible_attorney_id\n" +
			"status: intake\n" +
			"subtype: subtype\n" +
			"vault:\n" +
			"  description: description\n" +
			"  enableGraph: true\n" +
			"  enableIndexing: true\n" +
			"  metadata:\n" +
			"    foo: bar\n" +
			"vault_id: vault_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1", "create",
		)
	})
}

func TestMattersV1Retrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1", "retrieve",
			"--id", "id",
		)
	})
}

func TestMattersV1Update(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1", "update",
			"--id", "id",
			"--archived-at", "'2019-12-27T18:11:19.117Z'",
			"--billing", "{foo: bar}",
			"--client-name", "client_name",
			"--client-party-id", "client_party_id",
			"--custom-fields", "{foo: bar}",
			"--description", "description",
			"--display-id", "display_id",
			"--important-dates", "{foo: bar}",
			"--jurisdiction", "{foo: bar}",
			"--matter-type", "matter_type",
			"--metadata", "{foo: bar}",
			"--practice-area", "practice_area",
			"--responsible-attorney-id", "responsible_attorney_id",
			"--status", "intake",
			"--subtype", "subtype",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived_at: '2019-12-27T18:11:19.117Z'\n" +
			"billing:\n" +
			"  foo: bar\n" +
			"client_name: client_name\n" +
			"client_party_id: client_party_id\n" +
			"custom_fields:\n" +
			"  foo: bar\n" +
			"description: description\n" +
			"display_id: display_id\n" +
			"important_dates:\n" +
			"  foo: bar\n" +
			"jurisdiction:\n" +
			"  foo: bar\n" +
			"matter_type: matter_type\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"practice_area: practice_area\n" +
			"responsible_attorney_id: responsible_attorney_id\n" +
			"status: intake\n" +
			"subtype: subtype\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1", "update",
			"--id", "id",
		)
	})
}

func TestMattersV1List(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1", "list",
			"--matter-type", "matter_type",
			"--practice-area", "practice_area",
			"--query", "query",
			"--status", "status",
		)
	})
}
