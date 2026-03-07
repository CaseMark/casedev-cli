// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultEventsSubscriptionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:events:subscriptions", "create",
			"--api-key", "string",
			"--id", "id",
			"--callback-url", "https://example.com",
			"--event-type", "string",
			"--object-id", "string",
			"--signing-secret", "signingSecret",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"callbackUrl: https://example.com\n" +
			"eventTypes:\n" +
			"  - string\n" +
			"objectIds:\n" +
			"  - string\n" +
			"signingSecret: signingSecret\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "vault:events:subscriptions", "create",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestVaultEventsSubscriptionsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:events:subscriptions", "update",
			"--api-key", "string",
			"--id", "id",
			"--subscription-id", "subscriptionId",
			"--callback-url", "https://example.com",
			"--clear-signing-secret=true",
			"--event-type", "string",
			"--is-active=true",
			"--object-id", "string",
			"--signing-secret", "signingSecret",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"callbackUrl: https://example.com\n" +
			"clearSigningSecret: true\n" +
			"eventTypes:\n" +
			"  - string\n" +
			"isActive: true\n" +
			"objectIds:\n" +
			"  - string\n" +
			"signingSecret: signingSecret\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "vault:events:subscriptions", "update",
			"--api-key", "string",
			"--id", "id",
			"--subscription-id", "subscriptionId",
		)
	})
}

func TestVaultEventsSubscriptionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:events:subscriptions", "list",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestVaultEventsSubscriptionsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:events:subscriptions", "delete",
			"--api-key", "string",
			"--id", "id",
			"--subscription-id", "subscriptionId",
		)
	})
}

func TestVaultEventsSubscriptionsTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "vault:events:subscriptions", "test",
			"--api-key", "string",
			"--id", "id",
			"--subscription-id", "subscriptionId",
			"--event-type", "eventType",
			"--object-id", "objectId",
			"--payload", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventType: eventType\n" +
			"objectId: objectId\n" +
			"payload:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "vault:events:subscriptions", "test",
			"--api-key", "string",
			"--id", "id",
			"--subscription-id", "subscriptionId",
		)
	})
}
