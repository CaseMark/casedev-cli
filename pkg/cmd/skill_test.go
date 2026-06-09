// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestSkillsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "create",
			"--content", "x",
			"--name", "x",
			"--file", "{content: content, path: path, contentType: contentType, metadata: {}, name: name, summary: summary, tags: [string]}",
			"--metadata", "{}",
			"--slug", "slug",
			"--summary", "summary",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(skillsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "create",
			"--content", "x",
			"--name", "x",
			"--file.content", "content",
			"--file.path", "path",
			"--file.content-type", "contentType",
			"--file.metadata", "{}",
			"--file.name", "name",
			"--file.summary", "summary",
			"--file.tags", "[string]",
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
			"files:\n" +
			"  - content: content\n" +
			"    path: path\n" +
			"    contentType: contentType\n" +
			"    metadata: {}\n" +
			"    name: name\n" +
			"    summary: summary\n" +
			"    tags:\n" +
			"      - string\n" +
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
			"--file", "[{content: content, path: path, contentType: contentType, metadata: {}, name: name, summary: summary, tags: [string]}]",
			"--metadata", "{}",
			"--name", "name",
			"--slug", "slug",
			"--summary", "summary",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(skillsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "update",
			"--slug", "slug",
			"--content", "content",
			"--file.content", "content",
			"--file.path", "path",
			"--file.content-type", "contentType",
			"--file.metadata", "{}",
			"--file.name", "name",
			"--file.summary", "summary",
			"--file.tags", "[string]",
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
			"files:\n" +
			"  - content: content\n" +
			"    path: path\n" +
			"    contentType: contentType\n" +
			"    metadata: {}\n" +
			"    name: name\n" +
			"    summary: summary\n" +
			"    tags:\n" +
			"      - string\n" +
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

func TestSkillsExport(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"skills", "export",
			"--slug", "slug",
			"--target", "target",
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
