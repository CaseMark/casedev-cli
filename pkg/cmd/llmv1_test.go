// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestLlmV1CreateEmbedding(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"llm:v1", "create-embedding",
		"--input", "string",
		"--model", "model",
		"--dimensions", "0",
		"--encoding-format", "float",
		"--user", "user",
	)
}

func TestLlmV1ListModels(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"llm:v1", "list-models",
	)
}
