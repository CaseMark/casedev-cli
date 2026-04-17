// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultObjectsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "retrieve",
			"--id", "id",
			"--object-id", "objectId",
		)
	})
}

func TestVaultObjectsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "update",
			"--id", "id",
			"--object-id", "objectId",
			"--filename", "deposition-smith-2024.pdf",
			"--metadata", "{}",
			"--path", "/Discovery/Depositions",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"filename: deposition-smith-2024.pdf\n" +
			"metadata: {}\n" +
			"path: /Discovery/Depositions\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:objects", "update",
			"--id", "id",
			"--object-id", "objectId",
		)
	})
}

func TestVaultObjectsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "list",
			"--id", "id",
		)
	})
}

func TestVaultObjectsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "delete",
			"--id", "id",
			"--object-id", "objectId",
			"--force", "true",
		)
	})
}

func TestVaultObjectsCreatePresignedURL(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "create-presigned-url",
			"--id", "id",
			"--object-id", "objectId",
			"--content-type", "contentType",
			"--expires-in", "60",
			"--operation", "GET",
			"--size-bytes", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contentType: contentType\n" +
			"expiresIn: 60\n" +
			"operation: GET\n" +
			"sizeBytes: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:objects", "create-presigned-url",
			"--id", "id",
			"--object-id", "objectId",
		)
	})
}

func TestVaultObjectsDownload(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "download",
			"--id", "id",
			"--object-id", "objectId",
			"--output", "/dev/null",
		)
	})
}

func TestVaultObjectsGetChunks(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "get-chunks",
			"--id", "id",
			"--object-id", "objectId",
			"--end", "0",
			"--start", "0",
		)
	})
}

func TestVaultObjectsGetOcrWords(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "get-ocr-words",
			"--id", "id",
			"--object-id", "objectId",
			"--page", "0",
			"--word-end", "0",
			"--word-start", "0",
		)
	})
}

func TestVaultObjectsGetSummarizeJob(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "get-summarize-job",
			"--id", "id",
			"--object-id", "objectId",
			"--job-id", "jobId",
		)
	})
}

func TestVaultObjectsGetText(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "get-text",
			"--id", "id",
			"--object-id", "objectId",
		)
	})
}
