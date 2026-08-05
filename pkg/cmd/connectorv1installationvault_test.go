// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestConnectorsV1InstallationsVaultsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations:vaults", "list",
			"--id", "id",
		)
	})
}

func TestConnectorsV1InstallationsVaultsGrant(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations:vaults", "grant",
			"--id", "id",
			"--vault-id", "vaultId",
			"--can-manage=true",
			"--can-read=true",
			"--can-write=true",
			"--relationship", "owned",
			"--source", "provisioning",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"can_manage: true\n" +
			"can_read: true\n" +
			"can_write: true\n" +
			"relationship: owned\n" +
			"source: provisioning\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:installations:vaults", "grant",
			"--id", "id",
			"--vault-id", "vaultId",
		)
	})
}

func TestConnectorsV1InstallationsVaultsRevoke(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations:vaults", "revoke",
			"--id", "id",
			"--vault-id", "vaultId",
		)
	})
}
