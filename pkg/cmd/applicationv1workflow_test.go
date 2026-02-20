// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestApplicationsV1WorkflowsGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:workflows", "get-status",
		"--id", "id",
		"--project-id", "projectId",
	)
}
