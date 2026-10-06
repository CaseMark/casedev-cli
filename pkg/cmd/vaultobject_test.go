// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
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
			"--cursor", "cursor",
			"--file-origin", `{"provider":"clio"}`,
			"--include-totals=true",
			"--include-unconfirmed=true",
			"--limit", "1",
			"--query", "query",
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

func TestVaultObjectsAppend(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "append",
			"--id", "id",
			"--object-id", "objectId",
			"--append-object-id", "string",
			"--back-links=true",
			"--back-links-text", "backLinksText",
			"--bates", "{enabled: true, padTo: 0, prefix: prefix, start: 1, suffix: suffix}",
			"--client-reference", "clientReference",
			"--mode", "sync",
			"--rewrite-links=true",
			"--idempotency-key", "x",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(vaultObjectsAppend)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "append",
			"--id", "id",
			"--object-id", "objectId",
			"--append-object-id", "string",
			"--back-links=true",
			"--back-links-text", "backLinksText",
			"--bates.enabled=true",
			"--bates.pad-to", "0",
			"--bates.prefix", "prefix",
			"--bates.start", "1",
			"--bates.suffix", "suffix",
			"--client-reference", "clientReference",
			"--mode", "sync",
			"--rewrite-links=true",
			"--idempotency-key", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"appendObjectIds:\n" +
			"  - string\n" +
			"backLinks: true\n" +
			"backLinksText: backLinksText\n" +
			"bates:\n" +
			"  enabled: true\n" +
			"  padTo: 0\n" +
			"  prefix: prefix\n" +
			"  start: 1\n" +
			"  suffix: suffix\n" +
			"clientReference: clientReference\n" +
			"mode: sync\n" +
			"rewriteLinks: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:objects", "append",
			"--id", "id",
			"--object-id", "objectId",
			"--idempotency-key", "x",
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

func TestVaultObjectsGetPages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "get-pages",
			"--id", "id",
			"--object-id", "objectId",
			"--end", "0",
			"--start", "0",
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

func TestVaultObjectsMerge(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "merge",
			"--id", "id",
			"--filename", "filename",
			"--source-object-id", "string",
			"--source-rendition", "original",
			"--idempotency-key", "x",
			"--bates", "{padTo: 0, prefix: prefix, start: 1, suffix: suffix}",
			"--client-reference", "clientReference",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(vaultObjectsMerge)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "merge",
			"--id", "id",
			"--filename", "filename",
			"--source-object-id", "string",
			"--source-rendition", "original",
			"--idempotency-key", "x",
			"--bates.pad-to", "0",
			"--bates.prefix", "prefix",
			"--bates.start", "1",
			"--bates.suffix", "suffix",
			"--client-reference", "clientReference",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"filename: filename\n" +
			"sourceObjectIds:\n" +
			"  - string\n" +
			"sourceRendition: original\n" +
			"bates:\n" +
			"  padTo: 0\n" +
			"  prefix: prefix\n" +
			"  start: 1\n" +
			"  suffix: suffix\n" +
			"clientReference: clientReference\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:objects", "merge",
			"--id", "id",
			"--idempotency-key", "x",
		)
	})
}

func TestVaultObjectsMove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:objects", "move",
			"--id", "id",
			"--destination-vault-id", "destinationVaultId",
			"--mode", "move",
			"--object-id", "string",
			"--idempotency-key", "Idempotency-Key",
			"--path", "path",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destinationVaultId: destinationVaultId\n" +
			"mode: move\n" +
			"objectIds:\n" +
			"  - string\n" +
			"path: path\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:objects", "move",
			"--id", "id",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
