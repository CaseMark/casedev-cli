// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestPrivilegeV1Detect(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"privilege:v1", "detect",
		"--category", "attorney_client",
		"--content", "content",
		"--document-id", "document_id",
		"--include-rationale=true",
		"--jurisdiction", "US-Federal",
		"--model", "model",
		"--vault-id", "vault_id",
	)
}
