// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1GetPricing(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1", "get-pricing",
			"--api-key", "string",
		)
	})
}

func TestComputeV1GetUsage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "compute:v1", "get-usage",
			"--api-key", "string",
			"--month", "3",
			"--year", "2024",
		)
	})
}
