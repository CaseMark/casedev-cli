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
			t,
			"--api-key", "string",
			"vault:multipart", "abort",
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
			t, pipeData,
			"--api-key", "string",
			"vault:multipart", "abort",
			"--id", "id",
		)
	})
}

func TestVaultMultipartComplete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:multipart", "complete",
			"--id", "id",
			"--object-id", "objectId",
			"--part", "{etag: etag, partNumber: 1}",
			"--size-bytes", "1",
			"--upload-id", "uploadId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(vaultMultipartComplete)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:multipart", "complete",
			"--id", "id",
			"--object-id", "objectId",
			"--part.etag", "etag",
			"--part.part-number", "1",
			"--size-bytes", "1",
			"--upload-id", "uploadId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"objectId: objectId\n" +
			"parts:\n" +
			"  - etag: etag\n" +
			"    partNumber: 1\n" +
			"sizeBytes: 1\n" +
			"uploadId: uploadId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:multipart", "complete",
			"--id", "id",
		)
	})
}

func TestVaultMultipartGetPartURLs(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:multipart", "get-part-urls",
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
			t,
			"--api-key", "string",
			"vault:multipart", "get-part-urls",
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
			t, pipeData,
			"--api-key", "string",
			"vault:multipart", "get-part-urls",
			"--id", "id",
		)
	})
}

func TestVaultMultipartInit(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:multipart", "init",
			"--id", "id",
			"--content-type", "contentType",
			"--filename", "filename",
			"--size-bytes", "1",
			"--auto-index=true",
			"--is-ai-generated=true",
			"--metadata", "{}",
			"--part-size-bytes", "5242880",
			"--path", "path",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contentType: contentType\n" +
			"filename: filename\n" +
			"sizeBytes: 1\n" +
			"auto_index: true\n" +
			"is_ai_generated: true\n" +
			"metadata: {}\n" +
			"partSizeBytes: 5242880\n" +
			"path: path\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:multipart", "init",
			"--id", "id",
		)
	})
}
