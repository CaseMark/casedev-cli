// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestConnectorsV1LinksRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:links", "retrieve",
			"--id", "id",
		)
	})
}

func TestConnectorsV1LinksUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:links", "update",
			"--id", "id",
			"--mode", "once",
			"--policy", "{}",
			"--state", "paused",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"mode: once\n" +
			"policy: {}\n" +
			"state: paused\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:links", "update",
			"--id", "id",
		)
	})
}

func TestConnectorsV1LinksList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:links", "list",
			"--connection-id", "connection_id",
			"--direction", "import",
			"--mode", "once",
			"--pair-id", "pair_id",
			"--state", "ready",
			"--vault-id", "vault_id",
		)
	})
}

func TestConnectorsV1LinksDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:links", "delete",
			"--id", "id",
			"--vault-docs", "keep",
		)
	})
}

func TestConnectorsV1LinksListObjects(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:links", "list-objects",
			"--id", "id",
			"--cursor", "cursor",
			"--state", "pending",
		)
	})
}
