// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVoiceStreamingGetURL(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"voice:streaming", "get-url",
	)
}
