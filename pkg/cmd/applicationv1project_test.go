// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestApplicationsV1ProjectsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(applicationsV1ProjectsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"gitRepo: gitRepo\n" +
			"name: name\n" +
			"buildCommand: buildCommand\n" +
			"environmentVariables:\n" +
			"  - key: key\n" +
			"    target:\n" +
			"      - production\n" +
			"    value: value\n" +
			"    type: plain\n" +
			"framework: framework\n" +
			"gitBranch: gitBranch\n" +
			"installCommand: installCommand\n" +
			"outputDirectory: outputDirectory\n" +
			"rootDirectory: rootDirectory\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"applications:v1:projects", "create",
		)
	})
}

func TestApplicationsV1ProjectsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "retrieve",
			"--id", "id",
		)
	})
}

func TestApplicationsV1ProjectsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "list",
			"--enrich=true",
			"--limit", "0",
		)
	})
}

func TestApplicationsV1ProjectsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "delete",
			"--id", "id",
			"--delete-from-hosting=true",
		)
	})
}

func TestApplicationsV1ProjectsCreateDeployment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "create-deployment",
			"--id", "id",
			"--environment-variable", "{key: key, target: [production], value: value, type: plain}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(applicationsV1ProjectsCreateDeployment)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "create-deployment",
			"--id", "id",
			"--environment-variable.key", "key",
			"--environment-variable.target", "[production]",
			"--environment-variable.value", "value",
			"--environment-variable.type", "plain",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"environmentVariables:\n" +
			"  - key: key\n" +
			"    target:\n" +
			"      - production\n" +
			"    value: value\n" +
			"    type: plain\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"applications:v1:projects", "create-deployment",
			"--id", "id",
		)
	})
}

func TestApplicationsV1ProjectsCreateDomain(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "create-domain",
			"--id", "id",
			"--domain", "domain",
			"--git-branch", "gitBranch",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain: domain\n" +
			"gitBranch: gitBranch\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"applications:v1:projects", "create-domain",
			"--id", "id",
		)
	})
}

func TestApplicationsV1ProjectsCreateEnv(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "create-env",
			"--id", "id",
			"--key", "key",
			"--target", "production",
			"--value", "value",
			"--git-branch", "gitBranch",
			"--type", "plain",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"key: key\n" +
			"target:\n" +
			"  - production\n" +
			"value: value\n" +
			"gitBranch: gitBranch\n" +
			"type: plain\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"applications:v1:projects", "create-env",
			"--id", "id",
		)
	})
}

func TestApplicationsV1ProjectsDeleteDomain(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "delete-domain",
			"--id", "id",
			"--domain", "domain",
		)
	})
}

func TestApplicationsV1ProjectsDeleteEnv(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "delete-env",
			"--id", "id",
			"--env-id", "envId",
		)
	})
}

func TestApplicationsV1ProjectsGetRuntimeLogs(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "get-runtime-logs",
			"--id", "id",
			"--limit", "0",
		)
	})
}

func TestApplicationsV1ProjectsListDeployments(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "list-deployments",
			"--id", "id",
			"--limit", "0",
			"--state", "state",
			"--target", "production",
		)
	})
}

func TestApplicationsV1ProjectsListDomains(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "list-domains",
			"--id", "id",
		)
	})
}

func TestApplicationsV1ProjectsListEnv(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"applications:v1:projects", "list-env",
			"--id", "id",
			"--decrypt=true",
		)
	})
}
