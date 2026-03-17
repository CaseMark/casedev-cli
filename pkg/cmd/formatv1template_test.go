// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestFormatV1TemplatesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"format:v1:templates", "create",
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
			t, pipeData,
			"--api-key", "string",
			"format:v1:templates", "create",
		)
	})
}

func TestFormatV1TemplatesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"format:v1:templates", "retrieve",
			"--id", "id",
		)
	})
}

func TestFormatV1TemplatesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"format:v1:templates", "list",
			"--type", "type",
		)
	})
}
