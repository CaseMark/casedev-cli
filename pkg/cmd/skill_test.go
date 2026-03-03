// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestSkillsRead(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"skills", "read",
		"--slug", "slug",
	)
}

func TestSkillsResolve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"skills", "resolve",
		"--q", "q",
		"--limit", "1",
	)
}
