// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestLlmV1ChatCreateCompletion(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"llm:v1:chat", "create-completion",
		"--message", "{content: content, role: system}",
		"--casemark-show-reasoning=false",
		"--frequency-penalty", "0",
		"--max-tokens", "1000",
		"--model", "casemark/casemark-core-3",
		"--presence-penalty", "0",
		"--stream=false",
		"--temperature", "0.7",
		"--top-p", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(llmV1ChatCreateCompletion)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"llm:v1:chat", "create-completion",
		"--message.content", "content",
		"--message.role", "system",
		"--casemark-show-reasoning=false",
		"--frequency-penalty", "0",
		"--max-tokens", "1000",
		"--model", "casemark/casemark-core-3",
		"--presence-penalty", "0",
		"--stream=false",
		"--temperature", "0.7",
		"--top-p", "0",
	)
}
