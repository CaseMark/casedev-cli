// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestLegalV1Docket(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "docket",
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
}

func TestLegalV1Draft(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "draft",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(legalV1Draft)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "draft",
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
}

func TestLegalV1Find(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "find",
		"--api-key", "string",
		"--query", "xxx",
		"--jurisdiction", "jurisdiction",
		"--num-results", "1",
	)
}

func TestLegalV1GetCitations(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "get-citations",
		"--api-key", "string",
		"--text", "text",
	)
}

func TestLegalV1GetCitationsFromURL(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "get-citations-from-url",
		"--api-key", "string",
		"--url", "https://example.com",
	)
}

func TestLegalV1GetFullText(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "get-full-text",
		"--api-key", "string",
		"--url", "https://example.com",
		"--highlight-query", "highlightQuery",
		"--max-characters", "1000",
		"--summary-query", "summaryQuery",
	)
}

func TestLegalV1ListCourts(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "list-courts",
		"--api-key", "string",
		"--in-use-only=true",
		"--jurisdiction", "jurisdiction",
		"--limit", "1",
		"--offset", "0",
		"--query", "xx",
	)
}

func TestLegalV1ListJurisdictions(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "list-jurisdictions",
		"--api-key", "string",
		"--name", "xx",
	)
}

func TestLegalV1PatentSearch(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "patent-search",
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
}

func TestLegalV1Research(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "research",
		"--api-key", "string",
		"--query", "xxx",
		"--additional-query", "string",
		"--jurisdiction", "jurisdiction",
		"--num-results", "1",
	)
}

func TestLegalV1Similar(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "similar",
		"--api-key", "string",
		"--url", "https://example.com",
		"--jurisdiction", "jurisdiction",
		"--num-results", "1",
		"--start-published-date", "'2019-12-27'",
	)
}

func TestLegalV1TrademarkSearch(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "trademark-search",
		"--api-key", "string",
		"--registration-number", "registrationNumber",
		"--serial-number", "serialNumber",
	)
}

func TestLegalV1Verify(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "verify",
		"--api-key", "string",
		"--text", "text",
	)
}
