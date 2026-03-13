// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV1ChatFilesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat:files", "list",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAgentV1ChatFilesDownload(t *testing.T) {
	t.Skip("Mock server doesn't support application/octet-stream responses")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent:v1:chat:files", "download",
			"--api-key", "string",
			"--id", "id",
			"--file-path", "filePath",
			"--output", "/dev/null",
		)
	})
}
