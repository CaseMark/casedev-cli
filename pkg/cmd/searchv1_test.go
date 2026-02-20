// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestSearchV1Answer(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"search:v1", "answer",
		"--query", "query",
		"--exclude-domain", "string",
		"--include-domain", "string",
		"--max-tokens", "0",
		"--model", "model",
		"--num-results", "1",
		"--search-type", "auto",
		"--stream=true",
		"--temperature", "0",
		"--text=true",
		"--use-custom-llm=true",
	)
}

func TestSearchV1Contents(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"search:v1", "contents",
		"--url", "https://example.com",
		"--context", "context",
		"--extras", "{}",
		"--highlights=true",
		"--livecrawl=true",
		"--livecrawl-timeout", "0",
		"--subpages=true",
		"--subpage-target", "0",
		"--summary=true",
		"--text=true",
	)
}

func TestSearchV1Research(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"search:v1", "research",
		"--instructions", "instructions",
		"--model", "fast",
		"--output-schema", "{}",
		"--query", "query",
	)
}

func TestSearchV1RetrieveResearch(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"search:v1", "retrieve-research",
		"--id", "id",
		"--events", "events",
		"--stream=true",
	)
}

func TestSearchV1Search(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"search:v1", "search",
		"--query", "query",
		"--additional-query", "string",
		"--category", "category",
		"--contents", "contents",
		"--end-crawl-date", "2019-12-27",
		"--end-published-date", "2019-12-27",
		"--exclude-domain", "string",
		"--include-domain", "string",
		"--include-text=true",
		"--num-results", "1",
		"--start-crawl-date", "2019-12-27",
		"--start-published-date", "2019-12-27",
		"--type", "auto",
		"--user-location", "userLocation",
	)
}

func TestSearchV1Similar(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"search:v1", "similar",
		"--url", "https://example.com",
		"--contents", "contents",
		"--end-crawl-date", "2019-12-27",
		"--end-published-date", "2019-12-27",
		"--exclude-domain", "string",
		"--include-domain", "string",
		"--include-text=true",
		"--num-results", "1",
		"--start-crawl-date", "2019-12-27",
		"--start-published-date", "2019-12-27",
	)
}
