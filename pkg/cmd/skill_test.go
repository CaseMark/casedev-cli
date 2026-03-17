// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestSkillsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "create",
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
			t, pipeData,
			"--api-key", "string",
			"skills", "create",
		)
	})
}

func TestSkillsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "update",
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
			t, pipeData,
			"--api-key", "string",
			"skills", "update",
			"--slug", "slug",
		)
	})
}

func TestSkillsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "delete",
			"--slug", "slug",
		)
	})
}

func TestSkillsRead(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "read",
			"--slug", "slug",
		)
	})
}

func TestSkillsResolve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "resolve",
			"--q", "q",
			"--limit", "1",
		)
	})
}
