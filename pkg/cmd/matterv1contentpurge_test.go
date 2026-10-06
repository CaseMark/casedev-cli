// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1ContentPurgesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:content-purges", "create",
			"--id", "id",
			"--request-id", "request_id",
			"--object-id", "string",
			"--session-id", "string",
			"--transcription-id", "string",
			"--work-item-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"request_id: request_id\n" +
			"object_ids:\n" +
			"  - string\n" +
			"session_ids:\n" +
			"  - string\n" +
			"transcription_ids:\n" +
			"  - string\n" +
			"work_item_ids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:content-purges", "create",
			"--id", "id",
		)
	})
}

func TestMattersV1ContentPurgesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:content-purges", "retrieve",
			"--purge-id", "purgeId",
		)
	})
}
