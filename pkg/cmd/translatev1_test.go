// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestTranslateV1Detect(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"translate:v1", "detect",
		"--q", "string",
	)
}

func TestTranslateV1ListLanguages(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"translate:v1", "list-languages",
		"--model", "nmt",
		"--target", "target",
	)
}

func TestTranslateV1Translate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"translate:v1", "translate",
		"--q", "string",
		"--target", "es",
		"--format", "text",
		"--model", "nmt",
		"--source", "en",
	)
}
