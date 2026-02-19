// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultGraphragGetStats(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:graphrag", "get-stats",
		"--id", "id",
	)
}

func TestVaultGraphragInit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:graphrag", "init",
		"--id", "id",
	)
}

func TestVaultGraphragProcessObject(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:graphrag", "process-object",
		"--id", "id",
		"--object-id", "objectId",
	)
}
