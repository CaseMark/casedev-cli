// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestOperatorV1Create(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "operator:v1", "create",
			"--api-key", "string",
			"--name", "name",
			"--model", "model",
			"--size", "small",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"model: model\n" +
			"size: small\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "operator:v1", "create",
			"--api-key", "string",
		)
	})
}

func TestOperatorV1CreateChatCompletion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "operator:v1", "create-chat-completion",
			"--api-key", "string",
		)
	})
}

func TestOperatorV1CreateResponse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "operator:v1", "create-response",
			"--api-key", "string",
		)
	})
}

func TestOperatorV1GetStatus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "operator:v1", "get-status",
			"--api-key", "string",
		)
	})
}
