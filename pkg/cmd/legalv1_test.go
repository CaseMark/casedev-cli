// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestLegalV1Find(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "find",
		"--query", "xxx",
		"--jurisdiction", "jurisdiction",
		"--num-results", "1",
	)
}

func TestLegalV1GetCitations(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "get-citations",
		"--text", "text",
	)
}

func TestLegalV1GetCitationsFromURL(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "get-citations-from-url",
		"--url", "https://example.com",
	)
}

func TestLegalV1GetFullText(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "get-full-text",
		"--url", "https://example.com",
		"--highlight-query", "highlightQuery",
		"--max-characters", "1000",
		"--summary-query", "summaryQuery",
	)
}

func TestLegalV1ListJurisdictions(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "list-jurisdictions",
		"--name", "xx",
	)
}

func TestLegalV1PatentSearch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "patent-search",
		"--query", "x",
		"--application-status", "applicationStatus",
		"--application-type", "Utility",
		"--assignee", "assignee",
		"--filing-date-from", "2019-12-27",
		"--filing-date-to", "2019-12-27",
		"--grant-date-from", "2019-12-27",
		"--grant-date-to", "2019-12-27",
		"--inventor", "inventor",
		"--limit", "1",
		"--offset", "0",
		"--sort-by", "filingDate",
		"--sort-order", "asc",
	)
}

func TestLegalV1Research(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "research",
		"--query", "xxx",
		"--additional-query", "string",
		"--jurisdiction", "jurisdiction",
		"--num-results", "1",
	)
}

func TestLegalV1Similar(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "similar",
		"--url", "https://example.com",
		"--jurisdiction", "jurisdiction",
		"--num-results", "1",
		"--start-published-date", "2019-12-27",
	)
}

func TestLegalV1Verify(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legal:v1", "verify",
		"--text", "text",
	)
}
