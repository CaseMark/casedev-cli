// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1EnvironmentsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:environments", "create",
		"--name", "document-review-prod",
	)
}

func TestComputeV1EnvironmentsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:environments", "retrieve",
		"--name", "name",
	)
}

func TestComputeV1EnvironmentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:environments", "list",
	)
}

func TestComputeV1EnvironmentsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:environments", "delete",
		"--name", "litigation-processing",
	)
}

func TestComputeV1EnvironmentsSetDefault(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:v1:environments", "set-default",
		"--name", "prod-legal-docs",
	)
}
