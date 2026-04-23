// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestWebhooksV1EndpointsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "create",
			"--event-type-filter", "string",
			"--url", "https://example.com",
			"--description", "description",
			"--resource-scopes", "{matterIds: [string], vaultIds: [string]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webhooksV1EndpointsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "create",
			"--event-type-filter", "string",
			"--url", "https://example.com",
			"--description", "description",
			"--resource-scopes.matter-ids", "[string]",
			"--resource-scopes.vault-ids", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventTypeFilters:\n" +
			"  - string\n" +
			"url: https://example.com\n" +
			"description: description\n" +
			"resourceScopes:\n" +
			"  matterIds:\n" +
			"    - string\n" +
			"  vaultIds:\n" +
			"    - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks:v1:endpoints", "create",
		)
	})
}

func TestWebhooksV1EndpointsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "retrieve",
			"--id", "id",
		)
	})
}

func TestWebhooksV1EndpointsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "update",
			"--id", "id",
			"--description", "description",
			"--event-type-filter", "string",
			"--resource-scopes", "{matterIds: [string], vaultIds: [string]}",
			"--status", "active",
			"--url", "https://example.com",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webhooksV1EndpointsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "update",
			"--id", "id",
			"--description", "description",
			"--event-type-filter", "string",
			"--resource-scopes.matter-ids", "[string]",
			"--resource-scopes.vault-ids", "[string]",
			"--status", "active",
			"--url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"eventTypeFilters:\n" +
			"  - string\n" +
			"resourceScopes:\n" +
			"  matterIds:\n" +
			"    - string\n" +
			"  vaultIds:\n" +
			"    - string\n" +
			"status: active\n" +
			"url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks:v1:endpoints", "update",
			"--id", "id",
		)
	})
}

func TestWebhooksV1EndpointsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "list",
			"--limit", "1",
			"--status", "active",
		)
	})
}

func TestWebhooksV1EndpointsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "delete",
			"--id", "id",
		)
	})
}

func TestWebhooksV1EndpointsRotateSecret(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "rotate-secret",
			"--id", "id",
			"--previous-secret-expires-in-sec", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("previousSecretExpiresInSec: 0")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks:v1:endpoints", "rotate-secret",
			"--id", "id",
		)
	})
}

func TestWebhooksV1EndpointsTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:v1:endpoints", "test",
			"--id", "id",
			"--event-type", "eventType",
			"--payload", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventType: eventType\n" +
			"payload: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks:v1:endpoints", "test",
			"--id", "id",
		)
	})
}
