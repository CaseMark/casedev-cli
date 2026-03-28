// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1EventsSubscriptionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:events:subscriptions", "create",
			"--id", "id",
			"--callback-url", "https://example.com",
			"--event-type", "string",
			"--signing-secret", "signingSecret",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"callbackUrl: https://example.com\n" +
			"eventTypes:\n" +
			"  - string\n" +
			"signingSecret: signingSecret\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:events:subscriptions", "create",
			"--id", "id",
		)
	})
}

func TestMattersV1EventsSubscriptionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:events:subscriptions", "list",
			"--id", "id",
		)
	})
}

func TestMattersV1EventsSubscriptionsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:events:subscriptions", "delete",
			"--id", "id",
			"--subscription-id", "subscriptionId",
		)
	})
}
