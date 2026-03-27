// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestAgentV2ChatFilesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat:files", "list",
			"--id", "id",
		)
	})
}

func TestAgentV2ChatFilesDownload(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent:v2:chat:files", "download",
			"--id", "id",
			"--file-path", "filePath",
			"--output", "/dev/null",
		)
	})
}
