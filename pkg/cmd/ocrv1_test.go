// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestOcrV1Retrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ocr:v1", "retrieve",
			"--id", "id",
		)
	})
}

func TestOcrV1Download(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ocr:v1", "download",
			"--id", "id",
			"--type", "text",
			"--output", "/dev/null",
		)
	})
}

func TestOcrV1Process(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ocr:v1", "process",
			"--document-url", "https://example.com/contract.pdf",
			"--callback-url", "https://your-app.com/webhooks/ocr-complete",
			"--document-id", "contract-2024-001",
			"--engine", "doctr",
			"--features", "{embed: {}, forms: {foo: bar}, tables: {format: csv}}",
			"--result-bucket", "my-ocr-results",
			"--result-prefix", "ocr/2024/",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(ocrV1Process)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ocr:v1", "process",
			"--document-url", "https://example.com/contract.pdf",
			"--callback-url", "https://your-app.com/webhooks/ocr-complete",
			"--document-id", "contract-2024-001",
			"--engine", "doctr",
			"--features.embed", "{}",
			"--features.forms", "{foo: bar}",
			"--features.tables", "{format: csv}",
			"--result-bucket", "my-ocr-results",
			"--result-prefix", "ocr/2024/",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"document_url: https://example.com/contract.pdf\n" +
			"callback_url: https://your-app.com/webhooks/ocr-complete\n" +
			"document_id: contract-2024-001\n" +
			"engine: doctr\n" +
			"features:\n" +
			"  embed: {}\n" +
			"  forms:\n" +
			"    foo: bar\n" +
			"  tables:\n" +
			"    format: csv\n" +
			"result_bucket: my-ocr-results\n" +
			"result_prefix: ocr/2024/\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ocr:v1", "process",
		)
	})
}
