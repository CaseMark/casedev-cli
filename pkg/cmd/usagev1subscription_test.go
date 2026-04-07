// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestUsageV1SubscriptionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage:v1:subscriptions", "create",
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
			"usage:v1:subscriptions", "create",
		)
	})
}

func TestUsageV1SubscriptionsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage:v1:subscriptions", "update",
			"--subscription-id", "subscriptionId",
			"--callback-url", "https://example.com",
			"--clear-signing-secret=true",
			"--event-type", "string",
			"--is-active=true",
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
			"signingSecret: signingSecret\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"usage:v1:subscriptions", "update",
			"--subscription-id", "subscriptionId",
		)
	})
}

func TestUsageV1SubscriptionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage:v1:subscriptions", "list",
		)
	})
}

func TestUsageV1SubscriptionsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage:v1:subscriptions", "delete",
			"--subscription-id", "subscriptionId",
		)
	})
}

func TestUsageV1SubscriptionsTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage:v1:subscriptions", "test",
			"--subscription-id", "subscriptionId",
			"--event-type", "eventType",
			"--payload", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventType: eventType\n" +
			"payload:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"usage:v1:subscriptions", "test",
			"--subscription-id", "subscriptionId",
		)
	})
}
