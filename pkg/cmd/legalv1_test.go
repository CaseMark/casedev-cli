// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestLegalV1Docket(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "docket",
			"--api-key", "string",
			"--type", "search",
			"--court", "court",
			"--date-filed-after", "'2019-12-27'",
			"--date-filed-before", "'2019-12-27'",
			"--docket-id", "docketId",
			"--include-entries=true",
			"--limit", "1",
			"--live=true",
			"--offset", "0",
			"--query", "xx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: search\n" +
			"court: court\n" +
			"dateFiledAfter: '2019-12-27'\n" +
			"dateFiledBefore: '2019-12-27'\n" +
			"docketId: docketId\n" +
			"includeEntries: true\n" +
			"limit: 1\n" +
			"live: true\n" +
			"offset: 0\n" +
			"query: xx\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "docket",
			"--api-key", "string",
		)
	})
}

func TestLegalV1Draft(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "draft",
			"--api-key", "string",
			"--instructions", "xxxxxxxxxx",
			"--vault-id", "vault_id",
			"--citations=true",
			"--format", "format",
			"--length", "{target: 0, unit: words}",
			"--model", "model",
			"--object-id", "[string]",
			"--output-name", "output_name",
			"--output-type", "pdf",
			"--verified=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(legalV1Draft)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "draft",
			"--api-key", "string",
			"--instructions", "xxxxxxxxxx",
			"--vault-id", "vault_id",
			"--citations=true",
			"--format", "format",
			"--length.target", "0",
			"--length.unit", "words",
			"--model", "model",
			"--object-id", "[string]",
			"--output-name", "output_name",
			"--output-type", "pdf",
			"--verified=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: xxxxxxxxxx\n" +
			"vault_id: vault_id\n" +
			"citations: true\n" +
			"format: format\n" +
			"length:\n" +
			"  target: 0\n" +
			"  unit: words\n" +
			"model: model\n" +
			"object_ids:\n" +
			"  - string\n" +
			"output_name: output_name\n" +
			"output_type: pdf\n" +
			"verified: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "draft",
			"--api-key", "string",
		)
	})
}

func TestLegalV1Find(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "find",
			"--api-key", "string",
			"--query", "xxx",
			"--jurisdiction", "jurisdiction",
			"--num-results", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: xxx\n" +
			"jurisdiction: jurisdiction\n" +
			"numResults: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "find",
			"--api-key", "string",
		)
	})
}

func TestLegalV1GetCitations(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "get-citations",
			"--api-key", "string",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: text")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "get-citations",
			"--api-key", "string",
		)
	})
}

func TestLegalV1GetCitationsFromURL(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "get-citations-from-url",
			"--api-key", "string",
			"--url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("url: https://example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "get-citations-from-url",
			"--api-key", "string",
		)
	})
}

func TestLegalV1GetFullText(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "get-full-text",
			"--api-key", "string",
			"--url", "https://example.com",
			"--highlight-query", "highlightQuery",
			"--max-characters", "1000",
			"--summary-query", "summaryQuery",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"highlightQuery: highlightQuery\n" +
			"maxCharacters: 1000\n" +
			"summaryQuery: summaryQuery\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "get-full-text",
			"--api-key", "string",
		)
	})
}

func TestLegalV1ListCourts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "list-courts",
			"--api-key", "string",
			"--in-use-only=true",
			"--jurisdiction", "jurisdiction",
			"--limit", "1",
			"--offset", "0",
			"--query", "xx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inUseOnly: true\n" +
			"jurisdiction: jurisdiction\n" +
			"limit: 1\n" +
			"offset: 0\n" +
			"query: xx\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "list-courts",
			"--api-key", "string",
		)
	})
}

func TestLegalV1ListJurisdictions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "list-jurisdictions",
			"--api-key", "string",
			"--name", "xx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: xx")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "list-jurisdictions",
			"--api-key", "string",
		)
	})
}

func TestLegalV1PatentSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "patent-search",
			"--api-key", "string",
			"--query", "x",
			"--application-status", "applicationStatus",
			"--application-type", "Utility",
			"--assignee", "assignee",
			"--filing-date-from", "'2019-12-27'",
			"--filing-date-to", "'2019-12-27'",
			"--grant-date-from", "'2019-12-27'",
			"--grant-date-to", "'2019-12-27'",
			"--inventor", "inventor",
			"--limit", "1",
			"--offset", "0",
			"--sort-by", "filingDate",
			"--sort-order", "asc",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"applicationStatus: applicationStatus\n" +
			"applicationType: Utility\n" +
			"assignee: assignee\n" +
			"filingDateFrom: '2019-12-27'\n" +
			"filingDateTo: '2019-12-27'\n" +
			"grantDateFrom: '2019-12-27'\n" +
			"grantDateTo: '2019-12-27'\n" +
			"inventor: inventor\n" +
			"limit: 1\n" +
			"offset: 0\n" +
			"sortBy: filingDate\n" +
			"sortOrder: asc\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "patent-search",
			"--api-key", "string",
		)
	})
}

func TestLegalV1Research(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "research",
			"--api-key", "string",
			"--query", "xxx",
			"--additional-query", "string",
			"--jurisdiction", "jurisdiction",
			"--num-results", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: xxx\n" +
			"additionalQueries:\n" +
			"  - string\n" +
			"jurisdiction: jurisdiction\n" +
			"numResults: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "research",
			"--api-key", "string",
		)
	})
}

func TestLegalV1Similar(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "similar",
			"--api-key", "string",
			"--url", "https://example.com",
			"--jurisdiction", "jurisdiction",
			"--num-results", "1",
			"--start-published-date", "'2019-12-27'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"jurisdiction: jurisdiction\n" +
			"numResults: 1\n" +
			"startPublishedDate: '2019-12-27'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "similar",
			"--api-key", "string",
		)
	})
}

func TestLegalV1TrademarkSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "trademark-search",
			"--api-key", "string",
			"--registration-number", "registrationNumber",
			"--serial-number", "serialNumber",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"registrationNumber: registrationNumber\n" +
			"serialNumber: serialNumber\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "trademark-search",
			"--api-key", "string",
		)
	})
}

func TestLegalV1Verify(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legal:v1", "verify",
			"--api-key", "string",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: text")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legal:v1", "verify",
			"--api-key", "string",
		)
	})
}
