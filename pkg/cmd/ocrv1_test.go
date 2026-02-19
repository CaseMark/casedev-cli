// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestOcrV1Retrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ocr:v1", "retrieve",
		"--id", "id",
	)
}

func TestOcrV1Download(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ocr:v1", "download",
		"--id", "id",
		"--type", "text",
	)
}

func TestOcrV1Process(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ocr:v1", "process",
		"--document-url", "https://example.com/contract.pdf",
		"--callback-url", "https://your-app.com/webhooks/ocr-complete",
		"--document-id", "contract-2024-001",
		"--engine", "doctr",
		"--features", "{embed: {}, forms: {foo: bar}, tables: {format: csv}}",
		"--result-bucket", "my-ocr-results",
		"--result-prefix", "ocr/2024/",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(ocrV1Process)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
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
}
