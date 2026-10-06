// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestConnectorsV1ApplicationsKeysBind(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:applications:keys", "bind",
			"--id", "id",
			"--key-id", "keyId",
		)
	})
}

func TestConnectorsV1ApplicationsKeysRevoke(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors:v1:applications:keys", "revoke",
			"--id", "id",
			"--key-id", "keyId",
		)
	})
}
