// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVoiceV1ListVoices(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"voice:v1", "list-voices",
		"--category", "category",
		"--collection-id", "collection_id",
		"--include-total-count=true",
		"--next-page-token", "next_page_token",
		"--page-size", "1",
		"--search", "search",
		"--sort", "name",
		"--sort-direction", "asc",
		"--voice-type", "premade",
	)
}
