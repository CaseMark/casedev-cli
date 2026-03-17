// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestSearchV1Answer(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"excludeDomains:\n" +
			"  - string\n" +
			"includeDomains:\n" +
			"  - string\n" +
			"maxTokens: 0\n" +
			"model: model\n" +
			"numResults: 1\n" +
			"searchType: auto\n" +
			"stream: true\n" +
			"temperature: 0\n" +
			"text: true\n" +
			"useCustomLLM: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search:v1", "answer",
		)
	})
}

func TestSearchV1Contents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"urls:\n" +
			"  - https://example.com\n" +
			"context: context\n" +
			"extras: {}\n" +
			"highlights: true\n" +
			"livecrawl: true\n" +
			"livecrawlTimeout: 0\n" +
			"subpages: true\n" +
			"subpageTarget: 0\n" +
			"summary: true\n" +
			"text: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search:v1", "contents",
		)
	})
}

func TestSearchV1Research(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search:v1", "research",
			"--instructions", "instructions",
			"--model", "fast",
			"--output-schema", "{}",
			"--query", "query",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"model: fast\n" +
			"outputSchema: {}\n" +
			"query: query\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search:v1", "research",
		)
	})
}

func TestSearchV1RetrieveResearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search:v1", "retrieve-research",
			"--id", "id",
			"--events", "events",
			"--stream=true",
		)
	})
}

func TestSearchV1Search(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search:v1", "search",
			"--query", "query",
			"--additional-query", "string",
			"--category", "category",
			"--contents", "contents",
			"--end-crawl-date", "'2019-12-27'",
			"--end-published-date", "'2019-12-27'",
			"--exclude-domain", "string",
			"--include-domain", "string",
			"--include-text=true",
			"--num-results", "1",
			"--start-crawl-date", "'2019-12-27'",
			"--start-published-date", "'2019-12-27'",
			"--type", "auto",
			"--user-location", "userLocation",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"additionalQueries:\n" +
			"  - string\n" +
			"category: category\n" +
			"contents: contents\n" +
			"endCrawlDate: '2019-12-27'\n" +
			"endPublishedDate: '2019-12-27'\n" +
			"excludeDomains:\n" +
			"  - string\n" +
			"includeDomains:\n" +
			"  - string\n" +
			"includeText: true\n" +
			"numResults: 1\n" +
			"startCrawlDate: '2019-12-27'\n" +
			"startPublishedDate: '2019-12-27'\n" +
			"type: auto\n" +
			"userLocation: userLocation\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search:v1", "search",
		)
	})
}

func TestSearchV1Similar(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search:v1", "similar",
			"--url", "https://example.com",
			"--contents", "contents",
			"--end-crawl-date", "'2019-12-27'",
			"--end-published-date", "'2019-12-27'",
			"--exclude-domain", "string",
			"--include-domain", "string",
			"--include-text=true",
			"--num-results", "1",
			"--start-crawl-date", "'2019-12-27'",
			"--start-published-date", "'2019-12-27'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"contents: contents\n" +
			"endCrawlDate: '2019-12-27'\n" +
			"endPublishedDate: '2019-12-27'\n" +
			"excludeDomains:\n" +
			"  - string\n" +
			"includeDomains:\n" +
			"  - string\n" +
			"includeText: true\n" +
			"numResults: 1\n" +
			"startCrawlDate: '2019-12-27'\n" +
			"startPublishedDate: '2019-12-27'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search:v1", "similar",
		)
	})
}
