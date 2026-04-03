// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultMemoryCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:memory", "create",
			"--id", "id",
			"--content", "content",
			"--type", "fact",
			"--source", "source",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"type: fact\n" +
			"source: source\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:memory", "create",
			"--id", "id",
		)
	})
}

func TestVaultMemoryUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:memory", "update",
			"--id", "id",
			"--entry-id", "entryId",
			"--content", "content",
			"--source", "source",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"source: source\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:memory", "update",
			"--id", "id",
			"--entry-id", "entryId",
		)
	})
}

func TestVaultMemoryList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:memory", "list",
			"--id", "id",
		)
	})
}

func TestVaultMemoryDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:memory", "delete",
			"--id", "id",
			"--entry-id", "entryId",
		)
	})
}

func TestVaultMemorySearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:memory", "search",
			"--id", "id",
			"--query", "query",
			"--limit", "1",
			"--tag", "string",
			"--type", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"limit: 1\n" +
			"tags:\n" +
			"  - string\n" +
			"types:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:memory", "search",
			"--id", "id",
		)
	})
}
