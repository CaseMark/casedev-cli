// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestDatabaseV1ProjectsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "create",
			"--api-key", "string",
			"--name", "litigation-docs-db",
			"--description", "Production database for litigation document management",
			"--region", "aws-us-east-1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: litigation-docs-db\n" +
			"description: Production database for litigation document management\n" +
			"region: aws-us-east-1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "database:v1:projects", "create",
			"--api-key", "string",
		)
	})
}

func TestDatabaseV1ProjectsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestDatabaseV1ProjectsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "list",
			"--api-key", "string",
		)
	})
}

func TestDatabaseV1ProjectsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestDatabaseV1ProjectsCreateBranch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "create-branch",
			"--api-key", "string",
			"--id", "id",
			"--name", "staging",
			"--parent-branch-id", "branch_main_123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: staging\n" +
			"parentBranchId: branch_main_123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "database:v1:projects", "create-branch",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestDatabaseV1ProjectsGetConnection(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "get-connection",
			"--api-key", "string",
			"--id", "id",
			"--branch", "branch",
			"--pooled=true",
		)
	})
}

func TestDatabaseV1ProjectsListBranches(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "database:v1:projects", "list-branches",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
