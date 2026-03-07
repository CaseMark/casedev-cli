// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestVaultMultipartAbort(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:multipart", "abort",
			"--api-key", "string",
			"--id", "id",
			"--object-id", "objectId",
			"--upload-id", "uploadId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"objectId: objectId\n" +
			"uploadId: uploadId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "vault:multipart", "abort",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestVaultMultipartGetPartURLs(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:multipart", "get-part-urls",
			"--api-key", "string",
			"--id", "id",
			"--object-id", "objectId",
			"--part", "{partNumber: 1, sizeBytes: 1}",
			"--upload-id", "uploadId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(vaultMultipartGetPartURLs)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "vault:multipart", "get-part-urls",
			"--api-key", "string",
			"--id", "id",
			"--object-id", "objectId",
			"--part.part-number", "1",
			"--part.size-bytes", "1",
			"--upload-id", "uploadId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"objectId: objectId\n" +
			"parts:\n" +
			"  - partNumber: 1\n" +
			"    sizeBytes: 1\n" +
			"uploadId: uploadId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "vault:multipart", "get-part-urls",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
