// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestSkillsCustomList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "skills:custom", "list",
			"--api-key", "string",
			"--cursor", "cursor",
			"--limit", "1",
			"--tag", "tag",
		)
	})
}
