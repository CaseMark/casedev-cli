// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestMemoryV1Create(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "create",
			"--api-key", "string",
			"--message", "{content: content, role: user}",
			"--category", "category",
			"--extraction-prompt", "extraction_prompt",
			"--infer=true",
			"--metadata", "{foo: bar}",
			"--tag-1", "tag_1",
			"--tag-10", "tag_10",
			"--tag-11", "tag_11",
			"--tag-12", "tag_12",
			"--tag-2", "tag_2",
			"--tag-3", "tag_3",
			"--tag-4", "tag_4",
			"--tag-5", "tag_5",
			"--tag-6", "tag_6",
			"--tag-7", "tag_7",
			"--tag-8", "tag_8",
			"--tag-9", "tag_9",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoryV1Create)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "create",
			"--api-key", "string",
			"--message.content", "content",
			"--message.role", "user",
			"--category", "category",
			"--extraction-prompt", "extraction_prompt",
			"--infer=true",
			"--metadata", "{foo: bar}",
			"--tag-1", "tag_1",
			"--tag-10", "tag_10",
			"--tag-11", "tag_11",
			"--tag-12", "tag_12",
			"--tag-2", "tag_2",
			"--tag-3", "tag_3",
			"--tag-4", "tag_4",
			"--tag-5", "tag_5",
			"--tag-6", "tag_6",
			"--tag-7", "tag_7",
			"--tag-8", "tag_8",
			"--tag-9", "tag_9",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messages:\n" +
			"  - content: content\n" +
			"    role: user\n" +
			"category: category\n" +
			"extraction_prompt: extraction_prompt\n" +
			"infer: true\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"tag_1: tag_1\n" +
			"tag_10: tag_10\n" +
			"tag_11: tag_11\n" +
			"tag_12: tag_12\n" +
			"tag_2: tag_2\n" +
			"tag_3: tag_3\n" +
			"tag_4: tag_4\n" +
			"tag_5: tag_5\n" +
			"tag_6: tag_6\n" +
			"tag_7: tag_7\n" +
			"tag_8: tag_8\n" +
			"tag_9: tag_9\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "memory:v1", "create",
			"--api-key", "string",
		)
	})
}

func TestMemoryV1Retrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestMemoryV1List(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "list",
			"--api-key", "string",
			"--category", "category",
			"--limit", "0",
			"--offset", "0",
			"--tag-1", "tag_1",
			"--tag-10", "tag_10",
			"--tag-11", "tag_11",
			"--tag-12", "tag_12",
			"--tag-2", "tag_2",
			"--tag-3", "tag_3",
			"--tag-4", "tag_4",
			"--tag-5", "tag_5",
			"--tag-6", "tag_6",
			"--tag-7", "tag_7",
			"--tag-8", "tag_8",
			"--tag-9", "tag_9",
		)
	})
}

func TestMemoryV1Delete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestMemoryV1DeleteAll(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "delete-all",
			"--api-key", "string",
			"--tag-1", "tag_1",
			"--tag-10", "tag_10",
			"--tag-11", "tag_11",
			"--tag-12", "tag_12",
			"--tag-2", "tag_2",
			"--tag-3", "tag_3",
			"--tag-4", "tag_4",
			"--tag-5", "tag_5",
			"--tag-6", "tag_6",
			"--tag-7", "tag_7",
			"--tag-8", "tag_8",
			"--tag-9", "tag_9",
		)
	})
}

func TestMemoryV1Search(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "memory:v1", "search",
			"--api-key", "string",
			"--query", "query",
			"--category", "category",
			"--tag-1", "tag_1",
			"--tag-10", "tag_10",
			"--tag-11", "tag_11",
			"--tag-12", "tag_12",
			"--tag-2", "tag_2",
			"--tag-3", "tag_3",
			"--tag-4", "tag_4",
			"--tag-5", "tag_5",
			"--tag-6", "tag_6",
			"--tag-7", "tag_7",
			"--tag-8", "tag_8",
			"--tag-9", "tag_9",
			"--top-k", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"category: category\n" +
			"tag_1: tag_1\n" +
			"tag_10: tag_10\n" +
			"tag_11: tag_11\n" +
			"tag_12: tag_12\n" +
			"tag_2: tag_2\n" +
			"tag_3: tag_3\n" +
			"tag_4: tag_4\n" +
			"tag_5: tag_5\n" +
			"tag_6: tag_6\n" +
			"tag_7: tag_7\n" +
			"tag_8: tag_8\n" +
			"tag_9: tag_9\n" +
			"top_k: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "memory:v1", "search",
			"--api-key", "string",
		)
	})
}
