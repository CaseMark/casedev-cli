// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestTranslateV1Detect(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"translate:v1", "detect",
			"--q", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("q: string")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"translate:v1", "detect",
		)
	})
}

func TestTranslateV1ListLanguages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"translate:v1", "list-languages",
			"--model", "nmt",
			"--target", "target",
		)
	})
}

func TestTranslateV1Translate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"translate:v1", "translate",
			"--q", "string",
			"--target", "es",
			"--format", "text",
			"--model", "nmt",
			"--source", "en",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"q: string\n" +
			"target: es\n" +
			"format: text\n" +
			"model: nmt\n" +
			"source: en\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"translate:v1", "translate",
		)
	})
}
