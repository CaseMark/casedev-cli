// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestFormatV1TemplatesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "format:v1:templates", "create",
			"--api-key", "string",
			"--content", "content",
			"--name", "name",
			"--type", "caption",
			"--description", "description",
			"--styles", "{}",
			"--tag", "string",
			"--variable", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"name: name\n" +
			"type: caption\n" +
			"description: description\n" +
			"styles: {}\n" +
			"tags:\n" +
			"  - string\n" +
			"variables:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "format:v1:templates", "create",
			"--api-key", "string",
		)
	})
}

func TestFormatV1TemplatesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "format:v1:templates", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestFormatV1TemplatesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "format:v1:templates", "list",
			"--api-key", "string",
			"--type", "type",
		)
	})
}
