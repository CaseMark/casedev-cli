// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1SharesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:shares", "create",
			"--id", "id",
			"--target-org-id", "target_org_id",
			"--expires-at", "'2019-12-27T18:11:19.117Z'",
			"--permission", "read",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"target_org_id: target_org_id\n" +
			"expires_at: '2019-12-27T18:11:19.117Z'\n" +
			"permission: read\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:shares", "create",
			"--id", "id",
		)
	})
}

func TestMattersV1SharesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:shares", "list",
			"--id", "id",
		)
	})
}

func TestMattersV1SharesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:shares", "delete",
			"--id", "id",
			"--share-id", "shareId",
		)
	})
}
