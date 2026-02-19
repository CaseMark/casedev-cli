// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVoiceTranscriptionCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"voice:transcription", "create",
		"--audio-url", "audio_url",
		"--auto-highlights=true",
		"--boost-param", "low",
		"--content-safety-labels=true",
		"--format", "json",
		"--format-text=true",
		"--language-code", "language_code",
		"--language-detection=true",
		"--object-id", "object_id",
		"--punctuate=true",
		"--speaker-labels=true",
		"--speakers-expected", "0",
		"--speech-model", "string",
		"--vault-id", "vault_id",
		"--word-boost", "string",
	)
}

func TestVoiceTranscriptionRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"voice:transcription", "retrieve",
		"--id", "tr_abc123def456",
	)
}

func TestVoiceTranscriptionDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"voice:transcription", "delete",
		"--id", "id",
	)
}
