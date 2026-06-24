// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestMediaV1TranscriptsRetrieveCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"media:v1:transcripts:retrieve", "create",
			"--object-id", "object_id",
			"--vault-id", "vault_id",
			"--transcript", "{object_id: object_id, vault_id: vault_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mediaV1TranscriptsRetrieveCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"media:v1:transcripts:retrieve", "create",
			"--object-id", "object_id",
			"--vault-id", "vault_id",
			"--transcript.object-id", "object_id",
			"--transcript.vault-id", "vault_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"object_id: object_id\n" +
			"vault_id: vault_id\n" +
			"transcript:\n" +
			"  object_id: object_id\n" +
			"  vault_id: vault_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"media:v1:transcripts:retrieve", "create",
		)
	})
}
