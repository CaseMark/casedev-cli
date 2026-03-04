// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultObjectsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "retrieve",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
	)
}

func TestVaultObjectsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "update",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
		"--filename", "deposition-smith-2024.pdf",
		"--metadata", "{}",
		"--path", "/Discovery/Depositions",
	)
}

func TestVaultObjectsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "list",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestVaultObjectsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "delete",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
		"--force", "true",
	)
}

func TestVaultObjectsCreatePresignedURL(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "create-presigned-url",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
		"--content-type", "contentType",
		"--expires-in", "60",
		"--operation", "GET",
		"--size-bytes", "1",
	)
}

func TestVaultObjectsDownload(t *testing.T) {
	t.Skip("Mock server doesn't support application/octet-stream responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "download",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
		"--output", "/dev/null",
	)
}

func TestVaultObjectsGetOcrWords(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "get-ocr-words",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
		"--page", "0",
		"--word-end", "0",
		"--word-start", "0",
	)
}

func TestVaultObjectsGetSummarizeJob(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "get-summarize-job",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
		"--job-id", "jobId",
	)
}

func TestVaultObjectsGetText(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "get-text",
		"--api-key", "string",
		"--id", "id",
		"--object-id", "objectId",
	)
}
