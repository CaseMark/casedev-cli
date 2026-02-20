// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1InstancesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:instances", "create",
		"--instance-type", "gpu_1x_a10",
		"--name", "ocr-batch-job",
		"--region", "us-west-1",
		"--auto-shutdown-minutes", "120",
		"--vault-id", "vault_abc123",
	)
}

func TestComputeV1InstancesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:instances", "retrieve",
		"--id", "id",
	)
}

func TestComputeV1InstancesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:instances", "list",
	)
}

func TestComputeV1InstancesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:instances", "delete",
		"--id", "id",
	)
}
