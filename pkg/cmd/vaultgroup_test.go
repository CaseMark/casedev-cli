// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultGroupsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:groups", "create",
	)
}

func TestVaultGroupsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:groups", "update",
		"--group-id", "groupId",
	)
}

func TestVaultGroupsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:groups", "list",
	)
}

func TestVaultGroupsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:groups", "delete",
		"--group-id", "groupId",
	)
}
