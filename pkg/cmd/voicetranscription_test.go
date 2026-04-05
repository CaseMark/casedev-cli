// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVoiceTranscriptionCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audio_url: audio_url\n" +
			"auto_highlights: true\n" +
			"boost_param: low\n" +
			"content_safety_labels: true\n" +
			"format: json\n" +
			"format_text: true\n" +
			"language_code: language_code\n" +
			"language_detection: true\n" +
			"object_id: object_id\n" +
			"punctuate: true\n" +
			"speaker_labels: true\n" +
			"speakers_expected: 0\n" +
			"speech_models:\n" +
			"  - string\n" +
			"vault_id: vault_id\n" +
			"word_boost:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"voice:transcription", "create",
		)
	})
}

func TestVoiceTranscriptionRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice:transcription", "retrieve",
			"--id", "tr_abc123def456",
			"--include-text", "true",
		)
	})
}

func TestVoiceTranscriptionDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice:transcription", "delete",
			"--id", "id",
		)
	})
}
