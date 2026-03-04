// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestApplicationsV1DeploymentsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "create",
		"--api-key", "string",
		"--project-id", "projectId",
		"--ref", "ref",
		"--target", "production",
	)
}

func TestApplicationsV1DeploymentsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "retrieve",
		"--api-key", "string",
		"--id", "id",
		"--project-id", "projectId",
		"--include-logs=true",
	)
}

func TestApplicationsV1DeploymentsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "list",
		"--api-key", "string",
		"--project-id", "projectId",
		"--limit", "0",
		"--state", "state",
		"--target", "production",
	)
}

func TestApplicationsV1DeploymentsCancel(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "cancel",
		"--api-key", "string",
		"--id", "id",
		"--project-id", "projectId",
	)
}

func TestApplicationsV1DeploymentsCreateFromFiles(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "create-from-files",
		"--api-key", "string",
	)
}

func TestApplicationsV1DeploymentsGetLogs(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "get-logs",
		"--api-key", "string",
		"--id", "id",
		"--project-id", "projectId",
	)
}

func TestApplicationsV1DeploymentsGetStatus(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "get-status",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestApplicationsV1DeploymentsStream(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "stream",
		"--api-key", "string",
		"--id", "id",
		"--project-id", "projectId",
		"--start-index", "0",
	)
}
