// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVoiceBoostListExtract(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice:boost-list", "extract",
			"--category", "person",
			"--object-id", "string",
			"--text", "text",
			"--vault-id", "vault_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categories:\n" +
			"  - person\n" +
			"object_ids:\n" +
			"  - string\n" +
			"text: text\n" +
			"vault_id: vault_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"voice:boost-list", "extract",
		)
	})
}

func TestVoiceBoostListGenerate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice:boost-list", "generate",
			"--transcription-job-id", "transcription_job_id",
			"--category", "person",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"transcription_job_id: transcription_job_id\n" +
			"categories:\n" +
			"  - person\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"voice:boost-list", "generate",
		)
	})
}
