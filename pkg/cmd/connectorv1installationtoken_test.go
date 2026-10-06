// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestConnectorsV1InstallationsTokensCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations:tokens", "create",
			"--id", "id",
			"--scope", "read",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"scopes:\n" +
			"  - read\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:installations:tokens", "create",
			"--id", "id",
		)
	})
}

func TestConnectorsV1InstallationsTokensRevoke(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations:tokens", "revoke",
			"--id", "id",
			"--token-id", "tokenId",
		)
	})
}
