// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestLlmV1CreateEmbedding(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm:v1", "create-embedding",
			"--input", "string",
			"--model", "model",
			"--dimensions", "0",
			"--encoding-format", "float",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input: string\n" +
			"model: model\n" +
			"dimensions: 0\n" +
			"encoding_format: float\n" +
			"user: user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"llm:v1", "create-embedding",
		)
	})
}

func TestLlmV1ListModels(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm:v1", "list-models",
		)
	})
}
