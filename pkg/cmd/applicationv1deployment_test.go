// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestApplicationsV1DeploymentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "create",
			"--project-id", "projectId",
			"--ref", "ref",
			"--target", "production",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"projectId: projectId\n" +
			"ref: ref\n" +
			"target: production\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"applications:v1:deployments", "create",
		)
	})
}

func TestApplicationsV1DeploymentsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "retrieve",
			"--id", "id",
			"--project-id", "projectId",
			"--include-logs=true",
		)
	})
}

func TestApplicationsV1DeploymentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "list",
			"--project-id", "projectId",
			"--limit", "0",
			"--state", "state",
			"--target", "production",
		)
	})
}

func TestApplicationsV1DeploymentsCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "cancel",
			"--id", "id",
			"--project-id", "projectId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("projectId: projectId")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"applications:v1:deployments", "cancel",
			"--id", "id",
		)
	})
}

func TestApplicationsV1DeploymentsCreateFromFiles(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "create-from-files",
		)
	})
}

func TestApplicationsV1DeploymentsGetLogs(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "get-logs",
			"--id", "id",
			"--project-id", "projectId",
		)
	})
}

func TestApplicationsV1DeploymentsGetStatus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "get-status",
			"--id", "id",
		)
	})
}

func TestApplicationsV1DeploymentsStream(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:deployments", "stream",
			"--id", "id",
			"--project-id", "projectId",
			"--start-index", "0",
		)
	})
}
