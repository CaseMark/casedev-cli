// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestLincV1SessionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "create",
			"--document-template-slug", "[string]",
			"--idle-timeout-ms", "0",
			"--include-document-templates=true",
			"--instructions", "instructions",
			"--model", "model",
			"--scoped-api-key", "scopedApiKey",
			"--service-tier", "default",
			"--skill-slug", "[string]",
			"--title", "title",
			"--vault-id", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"documentTemplateSlugs:\n" +
			"  - string\n" +
			"idleTimeoutMs: 0\n" +
			"includeDocumentTemplates: true\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"scopedApiKey: scopedApiKey\n" +
			"serviceTier: default\n" +
			"skillSlugs:\n" +
			"  - string\n" +
			"title: title\n" +
			"vaultIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"linc:v1:sessions", "create",
		)
	})
}

func TestLincV1SessionsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "delete",
			"--id", "id",
		)
	})
}

func TestLincV1SessionsCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "cancel",
			"--id", "id",
			"--clear-queue=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("clearQueue: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"linc:v1:sessions", "cancel",
			"--id", "id",
		)
	})
}

func TestLincV1SessionsIngestEvents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "ingest-events",
			"--id", "id",
			"--frame", "{event: {foo: bar}, seq: 1, type: type}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(lincV1SessionsIngestEvents)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "ingest-events",
			"--id", "id",
			"--frame.event", "{foo: bar}",
			"--frame.seq", "1",
			"--frame.type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"frames:\n" +
			"  - event:\n" +
			"      foo: bar\n" +
			"    seq: 1\n" +
			"    type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"linc:v1:sessions", "ingest-events",
			"--id", "id",
		)
	})
}

func TestLincV1SessionsRetrieveEvents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "retrieve-events",
			"--id", "id",
			"--after-seq", "0",
			"--cursor", "0",
			"--exclude-event-type", "string",
			"--limit", "1",
		)
	})
}

func TestLincV1SessionsRetrieveMessages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "retrieve-messages",
			"--id", "id",
			"--after-seq", "0",
			"--cursor", "0",
			"--limit", "1",
		)
	})
}

func TestLincV1SessionsRetrieveState(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "retrieve-state",
			"--id", "id",
		)
	})
}

func TestLincV1SessionsSendRpc(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"linc:v1:sessions", "send-rpc",
			"--id", "id",
			"--type", "type",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: type\n" +
			"id: id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"linc:v1:sessions", "send-rpc",
			"--id", "id",
		)
	})
}
