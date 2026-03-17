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

var legalV1Docket = cli.Command{
	Name:    "docket",
	Usage:   "Search federal court dockets or retrieve a specific docket with optional filing\nentries. Use legal.listCourts() to resolve court slugs for filtering.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Search dockets or look up a docket by ID",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "acknowledge-pacer-fees",
			Usage:    "Required when live: true. Acknowledges that PACER fees (up to $3.00 per docket) plus a $0.05 service fee will be charged to your account.",
			Default:  false,
			BodyPath: "acknowledgePacerFees",
		},
		&requestflag.Flag[string]{
			Name:     "court",
			Usage:    `Optional court slug for filtering (e.g. "nysd", "ca9", "cafc"). Use legal.listCourts() to find slugs.`,
			BodyPath: "court",
		},
		&requestflag.Flag[any]{
			Name:     "date-filed-after",
			Usage:    "Optional lower bound for filing date (YYYY-MM-DD)",
			BodyPath: "dateFiledAfter",
		},
		&requestflag.Flag[any]{
			Name:     "date-filed-before",
			Usage:    "Optional upper bound for filing date (YYYY-MM-DD)",
			BodyPath: "dateFiledBefore",
		},
		&requestflag.Flag[string]{
			Name:     "docket-id",
			Usage:    "Docket ID (required for lookup)",
			BodyPath: "docketId",
		},
		&requestflag.Flag[bool]{
			Name:     "include-entries",
			Usage:    "Include docket entries/filings in lookup responses. Coming soon — currently returns 501. The parameter is accepted for forward compatibility.",
			Default:  false,
			BodyPath: "includeEntries",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Page size for search results or entry list (default 25 for search, 50 for lookup)",
			Default:  25,
			BodyPath: "limit",
		},
		&requestflag.Flag[bool]{
			Name:     "live",
			Usage:    `Trigger a live PACER fetch for dockets not yet in the RECAP archive. Requires acknowledgePacerFees: true. PACER charges up to $3.00 per docket sheet plus a $0.05 service fee. Only valid with type: "lookup".`,
			Default:  false,
			BodyPath: "live",
		},
		&requestflag.Flag[int64]{
			Name:     "offset",
			Usage:    "Offset for search results or entry list",
			Default:  0,
			BodyPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Case name or party name search query (required for search)",
			BodyPath: "query",
		},
	},
	Action:          handleLegalV1Docket,
	HideHelpCommand: true,
}

var legalV1Draft = requestflag.WithInnerFlags(cli.Command{
	Name:    "draft",
	Usage:   "Generate a legal document with structured inputs. Powered by an agent that\nhandles research, formatting, citation verification, and vault upload. Returns a\nrun ID for polling.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    `What to draft — the core task. E.g., "Motion to compel defendant to produce discovery responses"`,
			Required: true,
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Usage:    "Vault ID where the final document will be uploaded",
			Required: true,
			BodyPath: "vault_id",
		},
		&requestflag.Flag[bool]{
			Name:     "citations",
			Usage:    "Research and include legal citations",
			Default:  false,
			BodyPath: "citations",
		},
		&requestflag.Flag[any]{
			Name:     "format",
			Usage:    `Court or jurisdiction formatting hint. Triggers a legal-skills search. E.g., "California Superior Court", "SDNY", "federal pleading"`,
			BodyPath: "format",
		},
		&requestflag.Flag[any]{
			Name:     "length",
			Usage:    "Target document length",
			BodyPath: "length",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "LLM model override. Defaults to anthropic/claude-sonnet-4.6",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "object-id",
			Usage:    "Vault object IDs to use as source/reference documents",
			BodyPath: "object_ids",
		},
		&requestflag.Flag[any]{
			Name:     "output-name",
			Usage:    "Filename for the output document. Auto-generated if omitted.",
			BodyPath: "output_name",
		},
		&requestflag.Flag[string]{
			Name:     "output-type",
			Usage:    "Output file format",
			Default:  "md",
			BodyPath: "output_type",
		},
		&requestflag.Flag[bool]{
			Name:     "verified",
			Usage:    "Verify all citations in a loop — re-run verification and repair bad citations until they pass",
			Default:  false,
			BodyPath: "verified",
		},
	},
	Action:          handleLegalV1Draft,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"length": {
		&requestflag.InnerFlag[float64]{
			Name:       "length.target",
			Usage:      "Target value (e.g., 2000 words or 5 pages)",
			InnerField: "target",
		},
		&requestflag.InnerFlag[string]{
			Name:       "length.unit",
			Usage:      "Whether the target length is measured in words or pages",
			InnerField: "unit",
		},
	},
})

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

var legalV1ListCourts = cli.Command{
	Name:    "list-courts",
	Usage:   "Returns court IDs (slugs) and names for use with the docket search endpoint. Use\nthe returned court ID as the `court` parameter in legal.docket().",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "in-use-only",
			Usage:    "Only return courts with available docket data",
			Default:  true,
			BodyPath: "inUseOnly",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction",
			Usage:    "Optional jurisdiction code filter (e.g. FD for Federal District, F for all Federal, S for State)",
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum number of courts to return",
			Default:  25,
			BodyPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:     "offset",
			Usage:    "Number of courts to skip before returning results",
			Default:  0,
			BodyPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    `Search by court name or slug (e.g. "Northern District", "nysd", "ca9")`,
			BodyPath: "query",
		},
	},
	Action:          handleLegalV1ListCourts,
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

var legalV1SecFiling = cli.Command{
	Name:    "sec-filing",
	Usage:   "Search SEC EDGAR full-text filings via efts.sec.gov or fetch a filer's\nstructured filing history via data.sec.gov. Returns direct SEC archive URLs with\nfiling metadata and match snippets when available.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Run a full-text search or fetch a single entity filing history",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "cik",
			Usage:    "CIK for entity lookups. Accepts padded or unpadded digits.",
			BodyPath: "cik",
		},
		&requestflag.Flag[any]{
			Name:     "date-after",
			Usage:    "Optional lower filing date bound (YYYY-MM-DD)",
			BodyPath: "dateAfter",
		},
		&requestflag.Flag[any]{
			Name:     "date-before",
			Usage:    "Optional upper filing date bound (YYYY-MM-DD)",
			BodyPath: "dateBefore",
		},
		&requestflag.Flag[string]{
			Name:     "entity",
			Usage:    "Optional entity filter passed through to EDGAR full-text search",
			BodyPath: "entity",
		},
		&requestflag.Flag[[]string]{
			Name:     "form-type",
			Usage:    "Optional SEC form type filter such as 10-K, 10-Q, 8-K, or 4",
			BodyPath: "formTypes",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum filings to return",
			Default:  25,
			BodyPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:     "offset",
			Usage:    "Result offset for pagination",
			Default:  0,
			BodyPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Full-text SEC search query (required for type: search)",
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "ticker",
			Usage:    "Optional company ticker. Valid for both search and entity lookups.",
			BodyPath: "ticker",
		},
	},
	Action:          handleLegalV1SecFiling,
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

var legalV1TrademarkSearch = cli.Command{
	Name:    "trademark-search",
	Usage:   "Look up trademark status and details from the USPTO Trademark Status & Document\nRetrieval (TSDR) system. Supports lookup by serial number or registration\nnumber. Returns mark text, status, owner, goods/services, Nice classification,\nfiling/registration dates, and more.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "registration-number",
			Usage:    `USPTO registration number (e.g. "6123456"). Provide either serialNumber or registrationNumber.`,
			BodyPath: "registrationNumber",
		},
		&requestflag.Flag[string]{
			Name:     "serial-number",
			Usage:    `USPTO serial number (e.g. "97123456"). Provide either serialNumber or registrationNumber.`,
			BodyPath: "serialNumber",
		},
	},
	Action:          handleLegalV1TrademarkSearch,
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

func handleLegalV1Docket(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1DocketParams{}

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
	_, err = client.Legal.V1.Docket(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 docket", obj, format, transform)
}

func handleLegalV1Draft(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1DraftParams{}

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
	_, err = client.Legal.V1.Draft(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 draft", obj, format, transform)
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

func handleLegalV1ListCourts(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1ListCourtsParams{}

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
	_, err = client.Legal.V1.ListCourts(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 list-courts", obj, format, transform)
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

func handleLegalV1SecFiling(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1SecFilingParams{}

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
	_, err = client.Legal.V1.SecFiling(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 sec-filing", obj, format, transform)
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

func handleLegalV1TrademarkSearch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LegalV1TrademarkSearchParams{}

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
	_, err = client.Legal.V1.TrademarkSearch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legal:v1 trademark-search", obj, format, transform)
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
