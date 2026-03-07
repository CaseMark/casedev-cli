// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultGroupsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:groups", "create",
			"--api-key", "string",
		)
	})
}

func TestVaultGroupsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:groups", "update",
			"--api-key", "string",
			"--group-id", "groupId",
		)
	})
}

func TestVaultGroupsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:groups", "list",
			"--api-key", "string",
		)
	})
}

func TestVaultGroupsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:groups", "delete",
			"--api-key", "string",
			"--group-id", "groupId",
		)
	})
}
