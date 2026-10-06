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

var mattersV1Create = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new legal matter and optionally link an existing primary vault.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing",
			BodyPath: "billing",
		},
		&requestflag.Flag[string]{
			Name:     "client-name",
			BodyPath: "client_name",
		},
		&requestflag.Flag[*string]{
			Name:     "client-party-id",
			BodyPath: "client_party_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "custom_fields",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-id",
			BodyPath: "display_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "important-dates",
			BodyPath: "important_dates",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "jurisdiction",
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[string]{
			Name:     "matter-type",
			BodyPath: "matter_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "practice-area",
			BodyPath: "practice_area",
		},
		&requestflag.Flag[string]{
			Name:     "responsible-attorney-id",
			BodyPath: "responsible_attorney_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "intake", "open", "pending", "closed", "archived".`,
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "subtype",
			BodyPath: "subtype",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "vault",
			BodyPath: "vault",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			BodyPath: "vault_id",
		},
	},
	Action:          handleMattersV1Create,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"vault": {
		&requestflag.InnerFlag[string]{
			Name:       "vault.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "vault.enable-indexing",
			InnerField: "enableIndexing",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "vault.metadata",
			InnerField: "metadata",
		},
	},
})

var mattersV1Retrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single matter by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMattersV1Retrieve,
	HideHelpCommand: true,
}

var mattersV1Update = cli.Command{
	Name:    "update",
	Usage:   "Update mutable matter fields.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "archived-at",
			BodyPath: "archived_at",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing",
			BodyPath: "billing",
		},
		&requestflag.Flag[string]{
			Name:     "client-name",
			BodyPath: "client_name",
		},
		&requestflag.Flag[*string]{
			Name:     "client-party-id",
			BodyPath: "client_party_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "custom_fields",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-id",
			BodyPath: "display_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "important-dates",
			BodyPath: "important_dates",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "jurisdiction",
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[string]{
			Name:     "matter-type",
			BodyPath: "matter_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "practice-area",
			BodyPath: "practice_area",
		},
		&requestflag.Flag[string]{
			Name:     "responsible-attorney-id",
			BodyPath: "responsible_attorney_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "intake", "open", "pending", "closed", "archived".`,
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "subtype",
			BodyPath: "subtype",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			BodyPath: "title",
		},
	},
	Action:          handleMattersV1Update,
	HideHelpCommand: true,
}

var mattersV1List = cli.Command{
	Name:    "list",
	Usage:   "List matters for the authenticated organization, newest update first. Pagination\nis opt-in: pass `limit` (1-200) to receive a bounded page, then replay\n`pagination.next_cursor` as `?cursor=` while `pagination.has_more` is true.\nCursors are opaque and are only valid for the exact filter set they were issued\nunder. A request with neither `limit` nor `cursor` still returns every matter,\nand `pagination.limit` is null. That default will become a bounded page in a\nfuture release — paginate now to avoid the change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation cursor from `pagination.next_cursor` of the previous page. Must be replayed with the same filters that produced it.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Matters per page (1-200). Omit to receive every matter. Supplying a cursor without a limit uses 50.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "matter-type",
			QueryPath: "matter_type",
		},
		&requestflag.Flag[string]{
			Name:      "practice-area",
			QueryPath: "practice_area",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			QueryPath: "status",
		},
	},
	Action:          handleMattersV1List,
	HideHelpCommand: true,
}

var mattersV1Delete = cli.Command{
	Name:    "delete",
	Usage:   "Queues a durable, idempotent purge of a Matter and all linked live content. Use\nmatter purge webhooks for status changes; the inspection route is intended for\nmanual diagnostics only.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMattersV1Delete,
	HideHelpCommand: true,
}

func handleMattersV1Create(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := githubcomcasemarkcasedevgo.MatterV1NewParams{}

	return client.Matters.V1.New(ctx, params, options...)
}

func handleMattersV1Retrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Matters.V1.Get(ctx, cmd.Value("id").(string), options...)
}

func handleMattersV1Update(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomcasemarkcasedevgo.MatterV1UpdateParams{}

	return client.Matters.V1.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1List(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := githubcomcasemarkcasedevgo.MatterV1ListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Matters.V1.List(ctx, params, options...)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "matters:v1 list",
		Transform:      transform,
	})
}

func handleMattersV1Delete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Matters.V1.Delete(ctx, cmd.Value("id").(string), options...)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "matters:v1 delete",
		Transform:      transform,
	})
}
