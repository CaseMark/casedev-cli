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

var memoryV1Create = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Store memories from conversation messages. Automatically extracts facts and\nhandles deduplication.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Usage:    "Conversation messages to extract memories from",
			Required: true,
			BodyPath: "messages",
		},
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    `Custom category (e.g., "fact", "preference", "deadline")`,
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "extraction-prompt",
			Usage:    "Optional custom prompt for fact extraction",
			BodyPath: "extraction_prompt",
		},
		&requestflag.Flag[bool]{
			Name:     "infer",
			Usage:    "Whether to extract facts from messages (default: true)",
			Default:  true,
			BodyPath: "infer",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata (not indexed)",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "tag-1",
			Usage:    "Generic indexed filter field 1 (you decide what it means)",
			BodyPath: "tag_1",
		},
		&requestflag.Flag[string]{
			Name:     "tag-10",
			Usage:    "Generic indexed filter field 10",
			BodyPath: "tag_10",
		},
		&requestflag.Flag[string]{
			Name:     "tag-11",
			Usage:    "Generic indexed filter field 11",
			BodyPath: "tag_11",
		},
		&requestflag.Flag[string]{
			Name:     "tag-12",
			Usage:    "Generic indexed filter field 12",
			BodyPath: "tag_12",
		},
		&requestflag.Flag[string]{
			Name:     "tag-2",
			Usage:    "Generic indexed filter field 2",
			BodyPath: "tag_2",
		},
		&requestflag.Flag[string]{
			Name:     "tag-3",
			Usage:    "Generic indexed filter field 3",
			BodyPath: "tag_3",
		},
		&requestflag.Flag[string]{
			Name:     "tag-4",
			Usage:    "Generic indexed filter field 4",
			BodyPath: "tag_4",
		},
		&requestflag.Flag[string]{
			Name:     "tag-5",
			Usage:    "Generic indexed filter field 5",
			BodyPath: "tag_5",
		},
		&requestflag.Flag[string]{
			Name:     "tag-6",
			Usage:    "Generic indexed filter field 6",
			BodyPath: "tag_6",
		},
		&requestflag.Flag[string]{
			Name:     "tag-7",
			Usage:    "Generic indexed filter field 7",
			BodyPath: "tag_7",
		},
		&requestflag.Flag[string]{
			Name:     "tag-8",
			Usage:    "Generic indexed filter field 8",
			BodyPath: "tag_8",
		},
		&requestflag.Flag[string]{
			Name:     "tag-9",
			Usage:    "Generic indexed filter field 9",
			BodyPath: "tag_9",
		},
	},
	Action:          handleMemoryV1Create,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[string]{
			Name:       "message.content",
			Usage:      "Message content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.role",
			Usage:      "Message role",
			InnerField: "role",
		},
	},
})

var memoryV1Retrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single memory by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMemoryV1Retrieve,
	HideHelpCommand: true,
}

var memoryV1List = cli.Command{
	Name:    "list",
	Usage:   "List all memories with optional filtering by tags and category.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter by category",
			QueryPath: "category",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of results",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Pagination offset",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "tag-1",
			Usage:     "Filter by tag_1",
			QueryPath: "tag_1",
		},
		&requestflag.Flag[string]{
			Name:      "tag-10",
			Usage:     "Filter by tag_10",
			QueryPath: "tag_10",
		},
		&requestflag.Flag[string]{
			Name:      "tag-11",
			Usage:     "Filter by tag_11",
			QueryPath: "tag_11",
		},
		&requestflag.Flag[string]{
			Name:      "tag-12",
			Usage:     "Filter by tag_12",
			QueryPath: "tag_12",
		},
		&requestflag.Flag[string]{
			Name:      "tag-2",
			Usage:     "Filter by tag_2",
			QueryPath: "tag_2",
		},
		&requestflag.Flag[string]{
			Name:      "tag-3",
			Usage:     "Filter by tag_3",
			QueryPath: "tag_3",
		},
		&requestflag.Flag[string]{
			Name:      "tag-4",
			Usage:     "Filter by tag_4",
			QueryPath: "tag_4",
		},
		&requestflag.Flag[string]{
			Name:      "tag-5",
			Usage:     "Filter by tag_5",
			QueryPath: "tag_5",
		},
		&requestflag.Flag[string]{
			Name:      "tag-6",
			Usage:     "Filter by tag_6",
			QueryPath: "tag_6",
		},
		&requestflag.Flag[string]{
			Name:      "tag-7",
			Usage:     "Filter by tag_7",
			QueryPath: "tag_7",
		},
		&requestflag.Flag[string]{
			Name:      "tag-8",
			Usage:     "Filter by tag_8",
			QueryPath: "tag_8",
		},
		&requestflag.Flag[string]{
			Name:      "tag-9",
			Usage:     "Filter by tag_9",
			QueryPath: "tag_9",
		},
	},
	Action:          handleMemoryV1List,
	HideHelpCommand: true,
}

var memoryV1Delete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a single memory by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMemoryV1Delete,
	HideHelpCommand: true,
}

var memoryV1DeleteAll = cli.Command{
	Name:    "delete-all",
	Usage:   "Delete multiple memories matching tag filter criteria. CAUTION: This will delete\nall matching memories for your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "tag-1",
			Usage:     "Filter by tag_1",
			QueryPath: "tag_1",
		},
		&requestflag.Flag[string]{
			Name:      "tag-10",
			Usage:     "Filter by tag_10",
			QueryPath: "tag_10",
		},
		&requestflag.Flag[string]{
			Name:      "tag-11",
			Usage:     "Filter by tag_11",
			QueryPath: "tag_11",
		},
		&requestflag.Flag[string]{
			Name:      "tag-12",
			Usage:     "Filter by tag_12",
			QueryPath: "tag_12",
		},
		&requestflag.Flag[string]{
			Name:      "tag-2",
			Usage:     "Filter by tag_2",
			QueryPath: "tag_2",
		},
		&requestflag.Flag[string]{
			Name:      "tag-3",
			Usage:     "Filter by tag_3",
			QueryPath: "tag_3",
		},
		&requestflag.Flag[string]{
			Name:      "tag-4",
			Usage:     "Filter by tag_4",
			QueryPath: "tag_4",
		},
		&requestflag.Flag[string]{
			Name:      "tag-5",
			Usage:     "Filter by tag_5",
			QueryPath: "tag_5",
		},
		&requestflag.Flag[string]{
			Name:      "tag-6",
			Usage:     "Filter by tag_6",
			QueryPath: "tag_6",
		},
		&requestflag.Flag[string]{
			Name:      "tag-7",
			Usage:     "Filter by tag_7",
			QueryPath: "tag_7",
		},
		&requestflag.Flag[string]{
			Name:      "tag-8",
			Usage:     "Filter by tag_8",
			QueryPath: "tag_8",
		},
		&requestflag.Flag[string]{
			Name:      "tag-9",
			Usage:     "Filter by tag_9",
			QueryPath: "tag_9",
		},
	},
	Action:          handleMemoryV1DeleteAll,
	HideHelpCommand: true,
}

var memoryV1Search = cli.Command{
	Name:    "search",
	Usage:   "Search memories using semantic similarity. Filter by tag fields to narrow\nresults.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Search query for semantic matching",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "Filter by category",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "tag-1",
			Usage:    "Filter by tag_1",
			BodyPath: "tag_1",
		},
		&requestflag.Flag[string]{
			Name:     "tag-10",
			Usage:    "Filter by tag_10",
			BodyPath: "tag_10",
		},
		&requestflag.Flag[string]{
			Name:     "tag-11",
			Usage:    "Filter by tag_11",
			BodyPath: "tag_11",
		},
		&requestflag.Flag[string]{
			Name:     "tag-12",
			Usage:    "Filter by tag_12",
			BodyPath: "tag_12",
		},
		&requestflag.Flag[string]{
			Name:     "tag-2",
			Usage:    "Filter by tag_2",
			BodyPath: "tag_2",
		},
		&requestflag.Flag[string]{
			Name:     "tag-3",
			Usage:    "Filter by tag_3",
			BodyPath: "tag_3",
		},
		&requestflag.Flag[string]{
			Name:     "tag-4",
			Usage:    "Filter by tag_4",
			BodyPath: "tag_4",
		},
		&requestflag.Flag[string]{
			Name:     "tag-5",
			Usage:    "Filter by tag_5",
			BodyPath: "tag_5",
		},
		&requestflag.Flag[string]{
			Name:     "tag-6",
			Usage:    "Filter by tag_6",
			BodyPath: "tag_6",
		},
		&requestflag.Flag[string]{
			Name:     "tag-7",
			Usage:    "Filter by tag_7",
			BodyPath: "tag_7",
		},
		&requestflag.Flag[string]{
			Name:     "tag-8",
			Usage:    "Filter by tag_8",
			BodyPath: "tag_8",
		},
		&requestflag.Flag[string]{
			Name:     "tag-9",
			Usage:    "Filter by tag_9",
			BodyPath: "tag_9",
		},
		&requestflag.Flag[int64]{
			Name:     "top-k",
			Usage:    "Maximum number of results to return",
			Default:  10,
			BodyPath: "top_k",
		},
	},
	Action:          handleMemoryV1Search,
	HideHelpCommand: true,
}

func handleMemoryV1Create(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MemoryV1NewParams{}

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
	_, err = client.Memory.V1.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "memory:v1 create", obj, format, transform)
}

func handleMemoryV1Retrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Memory.V1.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "memory:v1 retrieve", obj, format, transform)
}

func handleMemoryV1List(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MemoryV1ListParams{}

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
	_, err = client.Memory.V1.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "memory:v1 list", obj, format, transform)
}

func handleMemoryV1Delete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Memory.V1.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "memory:v1 delete", obj, format, transform)
}

func handleMemoryV1DeleteAll(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MemoryV1DeleteAllParams{}

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
	_, err = client.Memory.V1.DeleteAll(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "memory:v1 delete-all", obj, format, transform)
}

func handleMemoryV1Search(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MemoryV1SearchParams{}

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
	_, err = client.Memory.V1.Search(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "memory:v1 search", obj, format, transform)
}
