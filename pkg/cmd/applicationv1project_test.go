// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestApplicationsV1ProjectsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "create",
		"--git-repo", "gitRepo",
		"--name", "name",
		"--build-command", "buildCommand",
		"--environment-variable", "{key: key, target: [production], value: value, type: plain}",
		"--framework", "framework",
		"--git-branch", "gitBranch",
		"--install-command", "installCommand",
		"--output-directory", "outputDirectory",
		"--root-directory", "rootDirectory",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(applicationsV1ProjectsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "create",
		"--git-repo", "gitRepo",
		"--name", "name",
		"--build-command", "buildCommand",
		"--environment-variable.key", "key",
		"--environment-variable.target", "[production]",
		"--environment-variable.value", "value",
		"--environment-variable.type", "plain",
		"--framework", "framework",
		"--git-branch", "gitBranch",
		"--install-command", "installCommand",
		"--output-directory", "outputDirectory",
		"--root-directory", "rootDirectory",
	)
}

func TestApplicationsV1ProjectsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "retrieve",
		"--id", "id",
	)
}

func TestApplicationsV1ProjectsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "list",
	)
}

func TestApplicationsV1ProjectsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "delete",
		"--id", "id",
		"--delete-from-hosting=true",
	)
}

func TestApplicationsV1ProjectsCreateDeployment(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "create-deployment",
		"--id", "id",
		"--environment-variable", "{key: key, target: [production], value: value, type: plain}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(applicationsV1ProjectsCreateDeployment)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "create-deployment",
		"--id", "id",
		"--environment-variable.key", "key",
		"--environment-variable.target", "[production]",
		"--environment-variable.value", "value",
		"--environment-variable.type", "plain",
	)
}

func TestApplicationsV1ProjectsCreateDomain(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "create-domain",
		"--id", "id",
		"--domain", "domain",
		"--git-branch", "gitBranch",
	)
}

func TestApplicationsV1ProjectsCreateEnv(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "create-env",
		"--id", "id",
		"--key", "key",
		"--target", "production",
		"--value", "value",
		"--git-branch", "gitBranch",
		"--type", "plain",
	)
}

func TestApplicationsV1ProjectsDeleteDomain(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "delete-domain",
		"--id", "id",
		"--domain", "domain",
	)
}

func TestApplicationsV1ProjectsDeleteEnv(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "delete-env",
		"--id", "id",
		"--env-id", "envId",
	)
}

func TestApplicationsV1ProjectsGetRuntimeLogs(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "get-runtime-logs",
		"--id", "id",
		"--limit", "0",
	)
}

func TestApplicationsV1ProjectsListDeployments(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "list-deployments",
		"--id", "id",
		"--limit", "0",
		"--state", "state",
		"--target", "production",
	)
}

func TestApplicationsV1ProjectsListDomains(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "list-domains",
		"--id", "id",
	)
}

func TestApplicationsV1ProjectsListEnv(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"applications:v1:projects", "list-env",
		"--id", "id",
		"--decrypt=true",
	)
}
