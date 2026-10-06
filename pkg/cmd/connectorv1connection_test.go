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
			"--provider", "box",
			"--return-url", "return_url",
			"--scope-tier", "box.readwrite",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider: box\n" +
			"return_url: return_url\n" +
			"scope_tier: box.readwrite\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:connections", "create",
			"--x-case-connector-subject", "x-case-connector-subject",
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
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})
}

func TestConnectorsV1ConnectionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "list",
			"--cursor", "cursor",
			"--limit", "1",
			"--provider", "provider",
			"--status", "pending",
			"--x-case-connector-subject", "x-case-connector-subject",
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
			"--x-case-connector-subject", "x-case-connector-subject",
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
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})
}

func TestConnectorsV1ConnectionsUpdateAll(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:connections", "update-all",
			"--confirm-organization-wide=true",
			"--enabled=true",
			"--provider", "provider",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"confirm_organization_wide: true\n" +
			"enabled: true\n" +
			"provider: provider\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:connections", "update-all",
		)
	})
}
