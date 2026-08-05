// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
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

func TestTranslateV1TranslateDocument(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"translate:v1", "translate-document",
			"--file", mocktest.TestFile(t, "Example data"),
			"--target", "es",
			"--source", "en",
			"--output", "/dev/null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"file: Example data\n" +
			"target: es\n" +
			"source: en\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"translate:v1", "translate-document",
			"--output", "/dev/null",
		)
	})
}
