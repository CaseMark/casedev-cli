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

var mattersV1PartiesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a reusable legal party for the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "address",
			BodyPath: "addresses",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "custom_fields",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			BodyPath: "notes",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "person", "organization".`,
			BodyPath: "type",
		},
	},
	Action:          handleMattersV1PartiesCreate,
	HideHelpCommand: true,
}

var mattersV1PartiesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a reusable legal party by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "party-id",
			Required:  true,
			PathParam: "partyId",
		},
	},
	Action:          handleMattersV1PartiesRetrieve,
	HideHelpCommand: true,
}

var mattersV1PartiesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a reusable legal party.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "party-id",
			Required:  true,
			PathParam: "partyId",
		},
	},
	Action:          handleMattersV1PartiesUpdate,
	HideHelpCommand: true,
}

var mattersV1PartiesList = cli.Command{
	Name:    "list",
	Usage:   "List reusable legal parties for the authenticated organization, newest update\nfirst. Pagination is opt-in: pass `limit` (1-200) to receive a bounded page,\nthen replay `pagination.next_cursor` as `?cursor=` while `pagination.has_more`\nis true. A request with neither `limit` nor `cursor` still returns every party,\nand `pagination.limit` is null. That default will become a bounded page in a\nfuture release — paginate now to avoid the change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation cursor from `pagination.next_cursor` of the previous page. Must be replayed with the same filters that produced it.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "email",
			QueryPath: "email",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Parties per page (1-200). Omit to receive every party. Supplying a cursor without a limit uses 50.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     `Allowed values: "person", "organization".`,
			QueryPath: "type",
		},
	},
	Action:          handleMattersV1PartiesList,
	HideHelpCommand: true,
}

func handleMattersV1PartiesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.MatterV1PartyNewParams{}

	return client.Matters.V1.Parties.New(ctx, params, options...)
}

func handleMattersV1PartiesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("party-id") && len(unusedArgs) > 0 {
		cmd.Set("party-id", unusedArgs[0])
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

	return client.Matters.V1.Parties.Get(ctx, cmd.Value("party-id").(string), options...)
}

func handleMattersV1PartiesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("party-id") && len(unusedArgs) > 0 {
		cmd.Set("party-id", unusedArgs[0])
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

	return client.Matters.V1.Parties.Update(ctx, cmd.Value("party-id").(string), options...)
}

func handleMattersV1PartiesList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.MatterV1PartyListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Matters.V1.Parties.List(ctx, params, options...)
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
		Title:          "matters:v1:parties list",
		Transform:      transform,
	})
}
