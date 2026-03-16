// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestSkillsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "skills", "create",
			"--api-key", "string",
			"--content", "x",
			"--name", "x",
			"--metadata", "{}",
			"--slug", "slug",
			"--summary", "summary",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: x\n" +
			"name: x\n" +
			"metadata: {}\n" +
			"slug: slug\n" +
			"summary: summary\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "skills", "create",
			"--api-key", "string",
		)
	})
}

func TestSkillsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "skills", "update",
			"--api-key", "string",
			"--slug", "slug",
			"--content", "content",
			"--metadata", "{}",
			"--name", "name",
			"--slug", "slug",
			"--summary", "summary",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"metadata: {}\n" +
			"name: name\n" +
			"slug: slug\n" +
			"summary: summary\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "skills", "update",
			"--api-key", "string",
			"--slug", "slug",
		)
	})
}

func TestSkillsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "skills", "delete",
			"--api-key", "string",
			"--slug", "slug",
		)
	})
}

func TestSkillsRead(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "skills", "read",
			"--api-key", "string",
			"--slug", "slug",
		)
	})
}

func TestSkillsResolve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "skills", "resolve",
			"--api-key", "string",
			"--q", "q",
			"--limit", "1",
		)
	})
}
