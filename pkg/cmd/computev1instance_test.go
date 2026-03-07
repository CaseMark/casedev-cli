// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1InstancesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:instances", "create",
			"--api-key", "string",
			"--instance-type", "gpu_1x_a10",
			"--name", "ocr-batch-job",
			"--region", "us-west-1",
			"--auto-shutdown-minutes", "120",
			"--vault-id", "vault_abc123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instanceType: gpu_1x_a10\n" +
			"name: ocr-batch-job\n" +
			"region: us-west-1\n" +
			"autoShutdownMinutes: 120\n" +
			"vaultIds:\n" +
			"  - vault_abc123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "compute:v1:instances", "create",
			"--api-key", "string",
		)
	})
}

func TestComputeV1InstancesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:instances", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestComputeV1InstancesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:instances", "list",
			"--api-key", "string",
		)
	})
}

func TestComputeV1InstancesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:instances", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
