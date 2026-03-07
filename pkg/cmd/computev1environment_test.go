// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1EnvironmentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:environments", "create",
			"--api-key", "string",
			"--name", "document-review-prod",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: document-review-prod")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "compute:v1:environments", "create",
			"--api-key", "string",
		)
	})
}

func TestComputeV1EnvironmentsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:environments", "retrieve",
			"--api-key", "string",
			"--name", "name",
		)
	})
}

func TestComputeV1EnvironmentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:environments", "list",
			"--api-key", "string",
		)
	})
}

func TestComputeV1EnvironmentsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:environments", "delete",
			"--api-key", "string",
			"--name", "litigation-processing",
		)
	})
}

func TestComputeV1EnvironmentsSetDefault(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1:environments", "set-default",
			"--api-key", "string",
			"--name", "prod-legal-docs",
		)
	})
}
