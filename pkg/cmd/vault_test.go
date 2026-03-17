// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestVaultCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "create",
			"--name", "Contract Review Archive",
			"--description", "Repository for all client contract reviews and analysis",
			"--enable-graph=true",
			"--enable-indexing=true",
			"--group-id", "grp_abc123",
			"--metadata", "{containsPHI: true, hipaaCompliant: true}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Contract Review Archive\n" +
			"description: Repository for all client contract reviews and analysis\n" +
			"enableGraph: true\n" +
			"enableIndexing: true\n" +
			"groupId: grp_abc123\n" +
			"metadata:\n" +
			"  containsPHI: true\n" +
			"  hipaaCompliant: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault", "create",
		)
	})
}

func TestVaultRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "retrieve",
			"--id", "vault_abc123",
		)
	})
}

func TestVaultUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "update",
			"--id", "id",
			"--description", "description",
			"--enable-graph=false",
			"--group-id", "groupId",
			"--name", "Updated Vault Name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"enableGraph: false\n" +
			"groupId: groupId\n" +
			"name: Updated Vault Name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault", "update",
			"--id", "id",
		)
	})
}

func TestVaultList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "list",
		)
	})
}

func TestVaultDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "delete",
			"--id", "id",
			"--async=true",
		)
	})
}

func TestVaultConfirmUpload(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "confirm-upload",
			"--id", "id",
			"--object-id", "objectId",
			"--size-bytes", "1",
			"--success=true",
			"--etag", "etag",
			"--error-code", "errorCode",
			"--error-message", "errorMessage",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sizeBytes: 1\n" +
			"success: true\n" +
			"etag: etag\n" +
			"errorCode: errorCode\n" +
			"errorMessage: errorMessage\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault", "confirm-upload",
			"--id", "id",
			"--object-id", "objectId",
		)
	})
}

func TestVaultIngest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "ingest",
			"--id", "id",
			"--object-id", "objectId",
		)
	})
}

func TestVaultSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "search",
			"--id", "id",
			"--query", "query",
			"--filters", "{object_id: string}",
			"--method", "vector",
			"--top-k", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(vaultSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "search",
			"--id", "id",
			"--query", "query",
			"--filters.object-id", "string",
			"--method", "vector",
			"--top-k", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"filters:\n" +
			"  object_id: string\n" +
			"method: vector\n" +
			"topK: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault", "search",
			"--id", "id",
		)
	})
}

func TestVaultUpload(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault", "upload",
			"--id", "id",
			"--content-type", "contentType",
			"--filename", "filename",
			"--auto-index=true",
			"--metadata", "{}",
			"--path", "path",
			"--size-bytes", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contentType: contentType\n" +
			"filename: filename\n" +
			"auto_index: true\n" +
			"metadata: {}\n" +
			"path: path\n" +
			"sizeBytes: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault", "upload",
			"--id", "id",
		)
	})
}
