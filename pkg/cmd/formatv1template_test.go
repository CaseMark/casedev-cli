// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestFormatV1TemplatesCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"format:v1:templates", "create",
		"--content", "content",
		"--name", "name",
		"--type", "caption",
		"--description", "description",
		"--styles", "{}",
		"--tag", "string",
		"--variable", "string",
	)
}

func TestFormatV1TemplatesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"format:v1:templates", "retrieve",
		"--id", "id",
	)
}

func TestFormatV1TemplatesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"format:v1:templates", "list",
		"--type", "type",
	)
}
