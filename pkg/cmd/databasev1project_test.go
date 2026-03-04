// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestDatabaseV1ProjectsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "create",
		"--api-key", "string",
		"--name", "litigation-docs-db",
		"--description", "Production database for litigation document management",
		"--region", "aws-us-east-1",
	)
}

func TestDatabaseV1ProjectsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "retrieve",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestDatabaseV1ProjectsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "list",
		"--api-key", "string",
	)
}

func TestDatabaseV1ProjectsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "delete",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestDatabaseV1ProjectsCreateBranch(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "create-branch",
		"--api-key", "string",
		"--id", "id",
		"--name", "staging",
		"--parent-branch-id", "branch_main_123",
	)
}

func TestDatabaseV1ProjectsGetConnection(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "get-connection",
		"--api-key", "string",
		"--id", "id",
		"--branch", "branch",
		"--pooled=true",
	)
}

func TestDatabaseV1ProjectsListBranches(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"database:v1:projects", "list-branches",
		"--api-key", "string",
		"--id", "id",
	)
}
