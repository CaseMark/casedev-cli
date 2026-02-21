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
		"--project-id", "projectId",
		"--ref", "ref",
		"--target", "production",
	)
}

func TestApplicationsV1DeploymentsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "retrieve",
		"--id", "id",
		"--project-id", "projectId",
		"--include-logs=true",
	)
}

func TestApplicationsV1DeploymentsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "list",
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
		"--id", "id",
		"--project-id", "projectId",
	)
}

func TestApplicationsV1DeploymentsCreateFromFiles(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "create-from-files",
	)
}

func TestApplicationsV1DeploymentsGetLogs(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "get-logs",
		"--id", "id",
		"--project-id", "projectId",
	)
}

func TestApplicationsV1DeploymentsGetStatus(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "get-status",
		"--id", "id",
	)
}

func TestApplicationsV1DeploymentsStream(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:deployments", "stream",
		"--id", "id",
		"--project-id", "projectId",
		"--start-index", "0",
	)
}
