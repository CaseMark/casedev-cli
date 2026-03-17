// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestLlmV1ChatCreateCompletion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm:v1:chat", "create-completion",
			"--message", "{content: content, role: system}",
			"--casemark-show-reasoning=false",
			"--frequency-penalty", "0",
			"--max-tokens", "1000",
			"--model", "casemark/casemark-core-6",
			"--presence-penalty", "0",
			"--stream=false",
			"--temperature", "0.7",
			"--top-p", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(llmV1ChatCreateCompletion)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm:v1:chat", "create-completion",
			"--message.content", "content",
			"--message.role", "system",
			"--casemark-show-reasoning=false",
			"--frequency-penalty", "0",
			"--max-tokens", "1000",
			"--model", "casemark/casemark-core-6",
			"--presence-penalty", "0",
			"--stream=false",
			"--temperature", "0.7",
			"--top-p", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messages:\n" +
			"  - content: content\n" +
			"    role: system\n" +
			"casemark_show_reasoning: false\n" +
			"frequency_penalty: 0\n" +
			"max_tokens: 1000\n" +
			"model: casemark/casemark-core-6\n" +
			"presence_penalty: 0\n" +
			"stream: false\n" +
			"temperature: 0.7\n" +
			"top_p: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"llm:v1:chat", "create-completion",
		)
	})
}
