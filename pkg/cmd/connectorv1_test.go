// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestConnectorsV1SyncLink(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1", "sync-link",
			"--connection-id", "connection_id",
			"--direction", "import",
			"--remote", "{folder_id: folder_id, container_id: container_id, path: path, resource_type: resource_type, site_id: site_id}",
			"--vault-id", "vault_id",
			"--export-destination", "{folder_id: folder_id, container_id: container_id, path: path, site_id: site_id}",
			"--matter-id", "matter_id",
			"--policy", "{collisions: version, deletes: mirror, filters: {exclude_file_ids: [string], exclude_folder_ids: [string], exclude_mime: [string], max_size_bytes: 0}}",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(connectorsV1SyncLink)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1", "sync-link",
			"--connection-id", "connection_id",
			"--direction", "import",
			"--remote.folder-id", "folder_id",
			"--remote.container-id", "container_id",
			"--remote.path", "path",
			"--remote.resource-type", "resource_type",
			"--remote.site-id", "site_id",
			"--vault-id", "vault_id",
			"--export-destination.folder-id", "folder_id",
			"--export-destination.container-id", "container_id",
			"--export-destination.path", "path",
			"--export-destination.site-id", "site_id",
			"--matter-id", "matter_id",
			"--policy.collisions", "version",
			"--policy.deletes", "mirror",
			"--policy.filters", "{exclude_file_ids: [string], exclude_folder_ids: [string], exclude_mime: [string], max_size_bytes: 0}",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connection_id: connection_id\n" +
			"direction: import\n" +
			"remote:\n" +
			"  folder_id: folder_id\n" +
			"  container_id: container_id\n" +
			"  path: path\n" +
			"  resource_type: resource_type\n" +
			"  site_id: site_id\n" +
			"vault_id: vault_id\n" +
			"export_destination:\n" +
			"  folder_id: folder_id\n" +
			"  container_id: container_id\n" +
			"  path: path\n" +
			"  site_id: site_id\n" +
			"matter_id: matter_id\n" +
			"policy:\n" +
			"  collisions: version\n" +
			"  deletes: mirror\n" +
			"  filters:\n" +
			"    exclude_file_ids:\n" +
			"      - string\n" +
			"    exclude_folder_ids:\n" +
			"      - string\n" +
			"    exclude_mime:\n" +
			"      - string\n" +
			"    max_size_bytes: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1", "sync-link",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})
}

func TestConnectorsV1Transfer(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1", "transfer",
			"--connection-id", "connection_id",
			"--direction", "import",
			"--remote", "{folder_id: folder_id, container_id: container_id, path: path, resource_type: resource_type, site_id: site_id}",
			"--vault-id", "vault_id",
			"--export-destination", "{folder_id: folder_id, container_id: container_id, path: path, site_id: site_id}",
			"--matter-id", "matter_id",
			"--policy", "{collisions: version, deletes: mirror, filters: {exclude_file_ids: [string], exclude_folder_ids: [string], exclude_mime: [string], max_size_bytes: 0}}",
			"--run-mode", "auto",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(connectorsV1Transfer)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1", "transfer",
			"--connection-id", "connection_id",
			"--direction", "import",
			"--remote.folder-id", "folder_id",
			"--remote.container-id", "container_id",
			"--remote.path", "path",
			"--remote.resource-type", "resource_type",
			"--remote.site-id", "site_id",
			"--vault-id", "vault_id",
			"--export-destination.folder-id", "folder_id",
			"--export-destination.container-id", "container_id",
			"--export-destination.path", "path",
			"--export-destination.site-id", "site_id",
			"--matter-id", "matter_id",
			"--policy.collisions", "version",
			"--policy.deletes", "mirror",
			"--policy.filters", "{exclude_file_ids: [string], exclude_folder_ids: [string], exclude_mime: [string], max_size_bytes: 0}",
			"--run-mode", "auto",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connection_id: connection_id\n" +
			"direction: import\n" +
			"remote:\n" +
			"  folder_id: folder_id\n" +
			"  container_id: container_id\n" +
			"  path: path\n" +
			"  resource_type: resource_type\n" +
			"  site_id: site_id\n" +
			"vault_id: vault_id\n" +
			"export_destination:\n" +
			"  folder_id: folder_id\n" +
			"  container_id: container_id\n" +
			"  path: path\n" +
			"  site_id: site_id\n" +
			"matter_id: matter_id\n" +
			"policy:\n" +
			"  collisions: version\n" +
			"  deletes: mirror\n" +
			"  filters:\n" +
			"    exclude_file_ids:\n" +
			"      - string\n" +
			"    exclude_folder_ids:\n" +
			"      - string\n" +
			"    exclude_mime:\n" +
			"      - string\n" +
			"    max_size_bytes: 0\n" +
			"run_mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1", "transfer",
			"--x-case-connector-subject", "x-case-connector-subject",
		)
	})
}
