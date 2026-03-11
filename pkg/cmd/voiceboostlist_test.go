// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVoiceBoostListExtract(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice:boost-list", "extract",
			"--api-key", "string",
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
			t, pipeData, "voice:boost-list", "extract",
			"--api-key", "string",
		)
	})
}

func TestVoiceBoostListGenerate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice:boost-list", "generate",
			"--api-key", "string",
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
			t, pipeData, "voice:boost-list", "generate",
			"--api-key", "string",
		)
	})
}
