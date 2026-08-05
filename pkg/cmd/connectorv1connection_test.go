// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestConnectorsV1ConnectionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "create",
			"--provider", "clio",
			"--return-url", "return_url",
			"--scope-tier", "clio.us",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider: clio\n" +
			"return_url: return_url\n" +
			"scope_tier: clio.us\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:connections", "create",
		)
	})
}

func TestConnectorsV1ConnectionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "retrieve",
			"--id", "id",
		)
	})
}

func TestConnectorsV1ConnectionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "list",
			"--provider", "provider",
			"--status", "pending",
		)
	})
}

func TestConnectorsV1ConnectionsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "delete",
			"--id", "id",
			"--purge=true",
		)
	})
}

func TestConnectorsV1ConnectionsBrowse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "browse",
			"--id", "id",
			"--container", "container",
			"--cursor", "cursor",
			"--page-size", "1000",
			"--parent", "parent",
			"--query", "query",
			"--site", "site",
		)
	})
}
