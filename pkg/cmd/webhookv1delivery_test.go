// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestWebhooksV1DeliveriesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:deliveries", "retrieve",
			"--id", "id",
		)
	})
}

func TestWebhooksV1DeliveriesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:deliveries", "list",
			"--endpoint-id", "endpoint_id",
			"--limit", "1",
			"--status", "pending",
		)
	})
}

func TestWebhooksV1DeliveriesReplay(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:deliveries", "replay",
			"--id", "id",
			"--payload", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("payload: {}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks:v1:deliveries", "replay",
			"--id", "id",
		)
	})
}
