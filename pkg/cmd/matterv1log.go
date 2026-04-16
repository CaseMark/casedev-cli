// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var mattersV1LogCreate = cli.Command{
	Name:    "create",
	Usage:   "Append a manual operational note or event to a matter log.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "summary",
			Required: true,
			BodyPath: "summary",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "details",
			BodyPath: "details",
		},
		&requestflag.Flag[string]{
			Name:     "event-type",
			BodyPath: "event_type",
		},
		&requestflag.Flag[string]{
			Name:     "work-item-id",
			BodyPath: "work_item_id",
		},
	},
	Action:          handleMattersV1LogCreate,
	HideHelpCommand: true,
}

var mattersV1LogList = cli.Command{
	Name:    "list",
	Usage:   "List the operational history for a matter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "actor-id",
			Usage:     "Filter by actor ID",
			QueryPath: "actor_id",
		},
		&requestflag.Flag[string]{
			Name:      "actor-type",
			Usage:     "Filter by actor type",
			QueryPath: "actor_type",
		},
		&requestflag.Flag[any]{
			Name:      "end-time",
			Usage:     "End of time range (ISO 8601)",
			QueryPath: "end_time",
		},
		&requestflag.Flag[string]{
			Name:      "event-type",
			Usage:     "Filter by exact event type",
			QueryPath: "event_type",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of log entries to return (max 200)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Number of log entries to skip for pagination",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[any]{
			Name:      "scope",
			Usage:     "Filter by scope: matter, work_item, execution, sharing, all",
			QueryPath: "scope",
		},
		&requestflag.Flag[any]{
			Name:      "start-time",
			Usage:     "Start of time range (ISO 8601)",
			QueryPath: "start_time",
		},
		&requestflag.Flag[string]{
			Name:      "work-item-id",
			Usage:     "Filter by work item ID",
			QueryPath: "work_item_id",
		},
	},
	Action:          handleMattersV1LogList,
	HideHelpCommand: true,
}

var mattersV1LogExport = cli.Command{
	Name:    "export",
	Usage:   "Bulk export matter log entries for audits, visibility, and eval pipelines.\nSupports json, csv, tsv, and jsonl. Limited to 10,000 entries per request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "actor-id",
			Usage:    "Filter by actor ID",
			BodyPath: "actor_id",
		},
		&requestflag.Flag[string]{
			Name:     "actor-type",
			Usage:    "Filter by actor type",
			BodyPath: "actor_type",
		},
		&requestflag.Flag[any]{
			Name:     "end-time",
			Usage:    "End of time range (ISO 8601)",
			BodyPath: "end_time",
		},
		&requestflag.Flag[string]{
			Name:     "event-type",
			Usage:    "Filter by exact event type",
			BodyPath: "event_type",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "Export format. Defaults to jsonl.",
			BodyPath: "format",
		},
		&requestflag.Flag[any]{
			Name:     "scope",
			Usage:    "Filter by scope: matter, work_item, execution, sharing, all",
			BodyPath: "scope",
		},
		&requestflag.Flag[any]{
			Name:     "start-time",
			Usage:    "Start of time range (ISO 8601)",
			BodyPath: "start_time",
		},
		&requestflag.Flag[string]{
			Name:     "work-item-id",
			Usage:    "Filter by work item ID",
			BodyPath: "work_item_id",
		},
	},
	Action:          handleMattersV1LogExport,
	HideHelpCommand: true,
}

func handleMattersV1LogCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1LogNewParams{}

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

	return client.Matters.V1.Log.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1LogList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1LogListParams{}

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

	return client.Matters.V1.Log.List(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1LogExport(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1LogExportParams{}

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
	_, err = client.Matters.V1.Log.Export(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "matters:v1:log export",
		Transform:      transform,
	})
}
