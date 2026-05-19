// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestAgentSkillsNamespacesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "create",
			"--namespace-id", "namespaceId",
			"--description", "description",
			"--label", "label",
			"--metadata", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"namespaceId: namespaceId\n" +
			"description: description\n" +
			"label: label\n" +
			"metadata: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:skills:namespaces", "create",
		)
	})
}

func TestAgentSkillsNamespacesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "retrieve",
			"--id", "id",
		)
	})
}

func TestAgentSkillsNamespacesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "list",
		)
	})
}

func TestAgentSkillsNamespacesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "delete",
			"--id", "id",
		)
	})
}

func TestAgentSkillsNamespacesPublish(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "publish",
			"--id", "id",
			"--file", "{content: content, encoding: utf8, path: path, contentType: contentType}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentSkillsNamespacesPublish)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "publish",
			"--id", "id",
			"--file.content", "content",
			"--file.encoding", "utf8",
			"--file.path", "path",
			"--file.content-type", "contentType",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"files:\n" +
			"  - content: content\n" +
			"    encoding: utf8\n" +
			"    path: path\n" +
			"    contentType: contentType\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent:skills:namespaces", "publish",
			"--id", "id",
		)
	})
}

func TestAgentSkillsNamespacesPull(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "pull",
			"--id", "id",
		)
	})
}

func TestAgentSkillsNamespacesRotateToken(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:skills:namespaces", "rotate-token",
			"--id", "id",
		)
	})
}
