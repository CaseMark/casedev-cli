// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultGraphragGetStats(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:graphrag", "get-stats",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestVaultGraphragInit(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:graphrag", "init",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestVaultGraphragProcessObject(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:graphrag", "process-object",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
	)
}
