// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1SecretsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:secrets", "create",
		"--api-key", "string",
		"--name", "name",
		"--description", "description",
		"--env", "env",
	)
}

func TestComputeV1SecretsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:secrets", "list",
		"--api-key", "string",
		"--env", "env",
	)
}

func TestComputeV1SecretsDeleteGroup(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:secrets", "delete-group",
		"--api-key", "string",
		"--group", "group",
		"--env", "env",
		"--key", "key",
	)
}

func TestComputeV1SecretsRetrieveGroup(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:secrets", "retrieve-group",
		"--api-key", "string",
		"--group", "group",
		"--env", "env",
	)
}

func TestComputeV1SecretsUpdateGroup(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:secrets", "update-group",
		"--api-key", "string",
		"--group", "litigation-apis",
		"--secrets", "{foo: string}",
		"--env", "env",
	)
}
