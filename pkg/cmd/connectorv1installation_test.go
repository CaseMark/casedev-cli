// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestConnectorsV1InstallationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations", "list",
			"--application", "application",
			"--external-tenant-id", "external_tenant_id",
		)
	})
}

func TestConnectorsV1InstallationsEnsure(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:installations", "ensure",
			"--application", "application",
			"--external-tenant-id", "external_tenant_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"application: application\n" +
			"external_tenant_id: external_tenant_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors:v1:installations", "ensure",
		)
	})
}
