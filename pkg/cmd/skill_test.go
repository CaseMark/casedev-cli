// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

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
