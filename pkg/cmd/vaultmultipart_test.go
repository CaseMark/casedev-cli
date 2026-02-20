// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestVaultMultipartAbort(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:multipart", "abort",
		"--id", "id",
		"--object-id", "objectId",
		"--upload-id", "uploadId",
	)
}

func TestVaultMultipartGetPartURLs(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:multipart", "get-part-urls",
		"--id", "id",
		"--object-id", "objectId",
		"--part", "{partNumber: 1, sizeBytes: 1}",
		"--upload-id", "uploadId",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(vaultMultipartGetPartURLs)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:multipart", "get-part-urls",
		"--id", "id",
		"--object-id", "objectId",
		"--part.part-number", "1",
		"--part.size-bytes", "1",
		"--upload-id", "uploadId",
	)
}
