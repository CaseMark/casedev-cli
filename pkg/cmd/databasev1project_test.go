// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestDatabaseV1ProjectsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "create",
		"--name", "litigation-docs-db",
		"--description", "Production database for litigation document management",
		"--region", "aws-us-east-1",
	)
}

func TestDatabaseV1ProjectsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "retrieve",
		"--id", "id",
	)
}

func TestDatabaseV1ProjectsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "list",
	)
}

func TestDatabaseV1ProjectsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "delete",
		"--id", "id",
	)
}

func TestDatabaseV1ProjectsCreateBranch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "create-branch",
		"--id", "id",
		"--name", "staging",
		"--parent-branch-id", "branch_main_123",
	)
}

func TestDatabaseV1ProjectsGetConnection(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "get-connection",
		"--id", "id",
		"--branch", "branch",
		"--pooled=true",
	)
}

func TestDatabaseV1ProjectsListBranches(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "list-branches",
		"--id", "id",
	)
}
