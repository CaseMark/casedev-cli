// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestVaultEventsSubscriptionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:events:subscriptions", "create",
		"--id", "id",
		"--callback-url", "https://example.com",
		"--event-type", "string",
		"--object-id", "string",
		"--signing-secret", "signingSecret",
	)
}

func TestVaultEventsSubscriptionsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:events:subscriptions", "update",
		"--id", "id",
		"--subscription-id", "subscriptionId",
		"--callback-url", "https://example.com",
		"--clear-signing-secret=true",
		"--event-type", "string",
		"--is-active=true",
		"--object-id", "string",
		"--signing-secret", "signingSecret",
	)
}

func TestVaultEventsSubscriptionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:events:subscriptions", "list",
		"--id", "id",
	)
}

func TestVaultEventsSubscriptionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:events:subscriptions", "delete",
		"--id", "id",
		"--subscription-id", "subscriptionId",
	)
}

func TestVaultEventsSubscriptionsTest(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault:events:subscriptions", "test",
		"--id", "id",
		"--subscription-id", "subscriptionId",
		"--event-type", "eventType",
		"--object-id", "objectId",
		"--payload", "{foo: bar}",
	)
}
