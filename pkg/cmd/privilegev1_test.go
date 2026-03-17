// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestPrivilegeV1Detect(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"privilege:v1", "detect",
			"--category", "attorney_client",
			"--content", "content",
			"--document-id", "document_id",
			"--include-rationale=true",
			"--jurisdiction", "US-Federal",
			"--model", "model",
			"--vault-id", "vault_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categories:\n" +
			"  - attorney_client\n" +
			"content: content\n" +
			"document_id: document_id\n" +
			"include_rationale: true\n" +
			"jurisdiction: US-Federal\n" +
			"model: model\n" +
			"vault_id: vault_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"privilege:v1", "detect",
		)
	})
}
