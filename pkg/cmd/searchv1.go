// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var searchV1Answer = cli.Command{
	Name:    "answer",
	Usage:   "Generate comprehensive answers to questions using web search results. Supports\ntwo modes: native provider answers or custom LLM-powered answers using\nCase.dev's AI gateway. Perfect for legal research, fact-checking, and gathering\nsupporting evidence for cases.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The question or topic to research and answer",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    "Exclude these domains from search",
			BodyPath: "excludeDomains",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    "Only search within these domains",
			BodyPath: "includeDomains",
		},
		&requestflag.Flag[int64]{
			Name:     "max-tokens",
			Usage:    "Maximum tokens for LLM response",
			Default:  2048,
			BodyPath: "maxTokens",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "LLM model to use when useCustomLLM is true",
			Default:  "gpt-4o",
			BodyPath: "model",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of search results to consider",
			Default:  10,
			BodyPath: "numResults",
		},
		&requestflag.Flag[string]{
			Name:     "search-type",
			Usage:    "Type of search to perform",
			Default:  "auto",
			BodyPath: "searchType",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "Stream the response (only for native provider answers)",
			Default:  false,
			BodyPath: "stream",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			Usage:    "LLM temperature for answer generation",
			Default:  0.3,
			BodyPath: "temperature",
		},
		&requestflag.Flag[bool]{
			Name:     "text",
			Usage:    "Include text content in response",
			Default:  true,
			BodyPath: "text",
		},
		&requestflag.Flag[bool]{
			Name:     "use-custom-llm",
			Usage:    "Use Case.dev LLM for answer generation instead of provider's native answer",
			Default:  false,
			BodyPath: "useCustomLLM",
		},
	},
	Action:          handleSearchV1Answer,
	HideHelpCommand: true,
}

var searchV1Contents = cli.Command{
	Name:    "contents",
	Usage:   "Scrapes and extracts text content from web pages, PDFs, and documents. Useful\nfor legal research, evidence collection, and document analysis. Supports live\ncrawling, subpage extraction, and content summarization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "url",
			Usage:    "Array of URLs to scrape and extract content from",
			Required: true,
			BodyPath: "urls",
		},
		&requestflag.Flag[string]{
			Name:     "context",
			Usage:    "Context to guide content extraction and summarization",
			BodyPath: "context",
		},
		&requestflag.Flag[any]{
			Name:     "extras",
			Usage:    "Additional extraction options",
			BodyPath: "extras",
		},
		&requestflag.Flag[bool]{
			Name:     "highlights",
			Usage:    "Whether to include content highlights",
			Default:  false,
			BodyPath: "highlights",
		},
		&requestflag.Flag[bool]{
			Name:     "livecrawl",
			Usage:    "Whether to perform live crawling for dynamic content",
			Default:  false,
			BodyPath: "livecrawl",
		},
		&requestflag.Flag[int64]{
			Name:     "livecrawl-timeout",
			Usage:    "Timeout in seconds for live crawling",
			Default:  30,
			BodyPath: "livecrawlTimeout",
		},
		&requestflag.Flag[bool]{
			Name:     "subpages",
			Usage:    "Whether to extract content from linked subpages",
			Default:  false,
			BodyPath: "subpages",
		},
		&requestflag.Flag[int64]{
			Name:     "subpage-target",
			Usage:    "Maximum number of subpages to crawl",
			Default:  5,
			BodyPath: "subpageTarget",
		},
		&requestflag.Flag[bool]{
			Name:     "summary",
			Usage:    "Whether to generate content summaries",
			Default:  false,
			BodyPath: "summary",
		},
		&requestflag.Flag[bool]{
			Name:     "text",
			Usage:    "Whether to extract text content",
			Default:  true,
			BodyPath: "text",
		},
	},
	Action:          handleSearchV1Contents,
	HideHelpCommand: true,
}

var searchV1Research = cli.Command{
	Name:    "research",
	Usage:   "Performs deep research by conducting multi-step analysis, gathering information\nfrom multiple sources, and providing comprehensive insights. Ideal for legal\nresearch, case analysis, and due diligence investigations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Research instructions or query",
			Required: true,
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Research quality level - fast (quick), normal (balanced), pro (comprehensive)",
			Default:  "normal",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			Usage:    "Optional JSON schema to structure the research output",
			BodyPath: "outputSchema",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Alias for instructions (for convenience)",
			BodyPath: "query",
		},
	},
	Action:          handleSearchV1Research,
	HideHelpCommand: true,
}

var searchV1RetrieveResearch = cli.Command{
	Name:    "retrieve-research",
	Usage:   "Retrieve the status and results of a deep research task by ID. Supports both\nstandard JSON responses and streaming for real-time updates as the research\nprogresses. Research tasks analyze topics comprehensively using web search and\nAI synthesis.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "events",
			Usage:     "Filter specific event types for streaming",
			QueryPath: "events",
		},
		&requestflag.Flag[bool]{
			Name:      "stream",
			Usage:     "Enable streaming for real-time updates",
			QueryPath: "stream",
		},
	},
	Action:          handleSearchV1RetrieveResearch,
	HideHelpCommand: true,
}

var searchV1Search = cli.Command{
	Name:    "search",
	Usage:   "Executes intelligent web search queries with advanced filtering and\ncustomization options. Ideal for legal research, case law discovery, and\ngathering supporting documentation for litigation or compliance matters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Primary search query",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[[]string]{
			Name:     "additional-query",
			Usage:    "Additional related search queries to enhance results",
			BodyPath: "additionalQueries",
		},
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "Category filter for search results",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "contents",
			Usage:    "Specific content type to search for",
			BodyPath: "contents",
		},
		&requestflag.Flag[any]{
			Name:     "end-crawl-date",
			Usage:    "End date for crawl date filtering",
			BodyPath: "endCrawlDate",
		},
		&requestflag.Flag[any]{
			Name:     "end-published-date",
			Usage:    "End date for published date filtering",
			BodyPath: "endPublishedDate",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    "Domains to exclude from search results",
			BodyPath: "excludeDomains",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    "Domains to include in search results",
			BodyPath: "includeDomains",
		},
		&requestflag.Flag[bool]{
			Name:     "include-text",
			Usage:    "Whether to include full text content in results",
			BodyPath: "includeText",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of search results to return",
			Default:  10,
			BodyPath: "numResults",
		},
		&requestflag.Flag[any]{
			Name:     "start-crawl-date",
			Usage:    "Start date for crawl date filtering",
			BodyPath: "startCrawlDate",
		},
		&requestflag.Flag[any]{
			Name:     "start-published-date",
			Usage:    "Start date for published date filtering",
			BodyPath: "startPublishedDate",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of search to perform",
			Default:  "auto",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "user-location",
			Usage:    "Geographic location for localized results",
			BodyPath: "userLocation",
		},
	},
	Action:          handleSearchV1Search,
	HideHelpCommand: true,
}

var searchV1Similar = cli.Command{
	Name:    "similar",
	Usage:   "Find web pages and documents similar to a given URL. Useful for legal research\nto discover related case law, statutes, or legal commentary that shares similar\nthemes or content structure.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL to find similar content for",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "contents",
			Usage:    "Additional content to consider for similarity matching",
			BodyPath: "contents",
		},
		&requestflag.Flag[any]{
			Name:     "end-crawl-date",
			Usage:    "Only include pages crawled before this date",
			BodyPath: "endCrawlDate",
		},
		&requestflag.Flag[any]{
			Name:     "end-published-date",
			Usage:    "Only include pages published before this date",
			BodyPath: "endPublishedDate",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    "Exclude results from these domains",
			BodyPath: "excludeDomains",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    "Only search within these domains",
			BodyPath: "includeDomains",
		},
		&requestflag.Flag[bool]{
			Name:     "include-text",
			Usage:    "Whether to include extracted text content in results",
			BodyPath: "includeText",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of similar results to return",
			Default:  10,
			BodyPath: "numResults",
		},
		&requestflag.Flag[any]{
			Name:     "start-crawl-date",
			Usage:    "Only include pages crawled after this date",
			BodyPath: "startCrawlDate",
		},
		&requestflag.Flag[any]{
			Name:     "start-published-date",
			Usage:    "Only include pages published after this date",
			BodyPath: "startPublishedDate",
		},
	},
	Action:          handleSearchV1Similar,
	HideHelpCommand: true,
}

func handleSearchV1Answer(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.SearchV1AnswerParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.V1.Answer(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search:v1 answer", obj, format, transform)
}

func handleSearchV1Contents(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.SearchV1ContentsParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.V1.Contents(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search:v1 contents", obj, format, transform)
}

func handleSearchV1Research(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.SearchV1ResearchParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.V1.Research(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search:v1 research", obj, format, transform)
}

func handleSearchV1RetrieveResearch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.SearchV1GetResearchParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.V1.GetResearch(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search:v1 retrieve-research", obj, format, transform)
}

func handleSearchV1Search(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.SearchV1SearchParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.V1.Search(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search:v1 search", obj, format, transform)
}

func handleSearchV1Similar(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.SearchV1SimilarParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.V1.Similar(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search:v1 similar", obj, format, transform)
}
