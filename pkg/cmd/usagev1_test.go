// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestUsageV1Retrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage:v1", "retrieve",
			"--granularity", "summary",
			"--period-end", "'2019-12-27T18:11:19.117Z'",
			"--period-start", "'2019-12-27T18:11:19.117Z'",
		)
	})
}
