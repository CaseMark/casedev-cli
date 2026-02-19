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

var legalV1Find = cli.Command{
	Name:    "find",
	Usage:   "Search for legal sources including cases, statutes, and regulations from\nauthoritative legal databases. Returns ranked candidates. Always verify with\nlegal.verify() before citing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    `Search query (e.g., "fair use copyright", "Miranda rights")`,
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction",
			Usage:    `Optional jurisdiction ID from resolveJurisdiction (e.g., "california", "us-federal")`,
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results 1-25 (default: 10)",
			Default:  10,
			BodyPath: "numResults",
		},
	},
	Action:          handleLegalV1Find,
	HideHelpCommand: true,
}

var legalV1GetCitations = cli.Command{
	Name:    "get-citations",
	Usage:   "Parses legal citations from text and returns structured Bluebook components\n(case name, reporter, volume, page, year, court). Accepts either a single\ncitation or a full text block.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    `Text containing citations to extract. Can be a single citation (e.g., "531 U.S. 98") or a full document with multiple citations.`,
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleLegalV1GetCitations,
	HideHelpCommand: true,
}

var legalV1GetCitationsFromURL = cli.Command{
	Name:    "get-citations-from-url",
	Usage:   "Extract all legal citations and references from a document URL. Returns\nstructured citation data including case citations, statute references, and\nregulatory citations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "URL of the legal document to extract citations from",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleLegalV1GetCitationsFromURL,
	HideHelpCommand: true,
}

var legalV1GetFullText = cli.Command{
	Name:    "get-full-text",
	Usage:   "Retrieve the full text content of a legal document. Use after verifying the\nsource with legal.verify(). Returns complete text with optional highlights and\nAI summary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "URL of the verified legal document",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "highlight-query",
			Usage:    `Optional query to extract relevant highlights (e.g., "What is the holding?")`,
			BodyPath: "highlightQuery",
		},
		&requestflag.Flag[int64]{
			Name:     "max-characters",
			Usage:    "Maximum characters to return (default: 10000, max: 50000)",
			Default:  10000,
			BodyPath: "maxCharacters",
		},
		&requestflag.Flag[string]{
			Name:     "summary-query",
			Usage:    `Optional query for generating a summary (e.g., "Summarize the key ruling")`,
			BodyPath: "summaryQuery",
		},
	},
	Action:          handleLegalV1GetFullText,
	HideHelpCommand: true,
}

var legalV1ListJurisdictions = cli.Command{
	Name:    "list-jurisdictions",
	Usage:   "Search for a jurisdiction by name. Returns matching jurisdictions with their IDs\nfor use in legal.find() and other legal research endpoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    `Jurisdiction name (e.g., "California", "US Federal", "NY")`,
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleLegalV1ListJurisdictions,
	HideHelpCommand: true,
}

var legalV1PatentSearch = cli.Command{
	Name:    "patent-search",
	Usage:   "Search the USPTO Open Data Portal for US patent applications and granted\npatents. Supports free-text queries, field-specific search, filters by\nassignee/inventor/status/type, date ranges, and pagination. Covers applications\nfiled on or after January 1, 2001. Data is refreshed daily.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    `Free-text search across all patent fields, or field-specific query (e.g. "applicationMetaData.patentNumber:11234567"). Supports AND, OR, NOT operators.`,
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "application-status",
			Usage:    `Filter by application status (e.g. "Patented Case", "Abandoned", "Pending")`,
			BodyPath: "applicationStatus",
		},
		&requestflag.Flag[string]{
			Name:     "application-type",
			Usage:    "Filter by application type",
			BodyPath: "applicationType",
		},
		&requestflag.Flag[string]{
			Name:     "assignee",
			Usage:    `Filter by assignee/owner name (e.g. "Google LLC")`,
			BodyPath: "assignee",
		},
		&requestflag.Flag[any]{
			Name:     "filing-date-from",
			Usage:    "Start of filing date range (YYYY-MM-DD)",
			BodyPath: "filingDateFrom",
		},
		&requestflag.Flag[any]{
			Name:     "filing-date-to",
			Usage:    "End of filing date range (YYYY-MM-DD)",
			BodyPath: "filingDateTo",
		},
		&requestflag.Flag[any]{
			Name:     "grant-date-from",
			Usage:    "Start of grant date range (YYYY-MM-DD)",
			BodyPath: "grantDateFrom",
		},
		&requestflag.Flag[any]{
			Name:     "grant-date-to",
			Usage:    "End of grant date range (YYYY-MM-DD)",
			BodyPath: "grantDateTo",
		},
		&requestflag.Flag[string]{
			Name:     "inventor",
			Usage:    "Filter by inventor name",
			BodyPath: "inventor",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Number of results to return (default 25, max 100)",
			Default:  25,
			BodyPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:     "offset",
			Usage:    "Starting position for pagination",
			Default:  0,
			BodyPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:     "sort-by",
			Usage:    "Field to sort results by",
			Default:  "filingDate",
			BodyPath: "sortBy",
		},
		&requestflag.Flag[string]{
			Name:     "sort-order",
			Usage:    "Sort order (default desc, newest first)",
			Default:  "desc",
			BodyPath: "sortOrder",
		},
	},
	Action:          handleLegalV1PatentSearch,
	HideHelpCommand: true,
}

var legalV1Research = cli.Command{
	Name:    "research",
	Usage:   "Perform comprehensive legal research with multiple query variations. Uses\nadvanced deep search to find relevant sources across different phrasings of the\nlegal issue.",
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
			Usage:    "Additional query variations to search (e.g., different phrasings of the legal issue)",
			BodyPath: "additionalQueries",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction",
			Usage:    "Optional jurisdiction ID from resolveJurisdiction",
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results 1-25 (default: 10)",
			Default:  10,
			BodyPath: "numResults",
		},
	},
	Action:          handleLegalV1Research,
	HideHelpCommand: true,
}

var legalV1Similar = cli.Command{
	Name:    "similar",
	Usage:   "Find cases and documents similar to a given legal source. Useful for finding\nciting cases, related precedents, or similar statutes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "URL of a legal document to find similar sources for",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction",
			Usage:    "Optional jurisdiction ID to filter results",
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results 1-25 (default: 10)",
			Default:  10,
			BodyPath: "numResults",
		},
		&requestflag.Flag[any]{
			Name:     "start-published-date",
			Usage:    `Optional ISO date to find only newer documents (e.g., "2020-01-01")`,
			BodyPath: "startPublishedDate",
		},
	},
	Action:          handleLegalV1Similar,
	HideHelpCommand: true,
}

var legalV1Verify = cli.Command{
	Name:    "verify",
	Usage:   "Validates legal citations against authoritative case law sources (CourtListener\ndatabase of ~10M cases). Returns verification status and case metadata for each\ncitation found in the input text. Accepts either a single citation or a full\ntext block containing multiple citations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    `Text containing citations to verify. Can be a single citation (e.g., "531 U.S. 98") or a full document with multiple citations.`,
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleLegalV1Verify,
	HideHelpCommand: true,
}

func handleLegalV1Find(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1FindParams{}

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
	_, err = client.Legal.V1.Find(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 find", obj, format, transform)
}

func handleLegalV1GetCitations(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1GetCitationsParams{}

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
	_, err = client.Legal.V1.GetCitations(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 get-citations", obj, format, transform)
}

func handleLegalV1GetCitationsFromURL(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1GetCitationsFromURLParams{}

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
	_, err = client.Legal.V1.GetCitationsFromURL(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 get-citations-from-url", obj, format, transform)
}

func handleLegalV1GetFullText(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1GetFullTextParams{}

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
	_, err = client.Legal.V1.GetFullText(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 get-full-text", obj, format, transform)
}

func handleLegalV1ListJurisdictions(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1ListJurisdictionsParams{}

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
	_, err = client.Legal.V1.ListJurisdictions(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 list-jurisdictions", obj, format, transform)
}

func handleLegalV1PatentSearch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1PatentSearchParams{}

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
	_, err = client.Legal.V1.PatentSearch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 patent-search", obj, format, transform)
}

func handleLegalV1Research(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1ResearchParams{}

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
	_, err = client.Legal.V1.Research(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 research", obj, format, transform)
}

func handleLegalV1Similar(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1SimilarParams{}

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
	_, err = client.Legal.V1.Similar(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 similar", obj, format, transform)
}

func handleLegalV1Verify(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1VerifyParams{}

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
	_, err = client.Legal.V1.Verify(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 verify", obj, format, transform)
}
