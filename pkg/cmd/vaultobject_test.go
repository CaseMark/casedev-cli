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
		"--id", "id",
		"--object-id", "objectId",
	)
}

func TestVaultObjectsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "update",
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
		"--id", "id",
	)
}

func TestVaultObjectsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "delete",
		"--id", "id",
		"--object-id", "objectId",
		"--force", "true",
	)
}

func TestVaultObjectsCreatePresignedURL(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "create-presigned-url",
		"--id", "id",
		"--object-id", "objectId",
		"--content-type", "contentType",
		"--expires-in", "60",
		"--operation", "GET",
		"--size-bytes", "1",
	)
}

func TestVaultObjectsDownload(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "download",
		"--id", "id",
		"--object-id", "objectId",
	)
}

func TestVaultObjectsGetOcrWords(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "get-ocr-words",
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
		"--id", "id",
		"--object-id", "objectId",
		"--job-id", "jobId",
	)
}

func TestVaultObjectsGetText(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:objects", "get-text",
		"--id", "id",
		"--object-id", "objectId",
	)
}
