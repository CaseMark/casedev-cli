// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
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
		&requestflag.Flag[any]{
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
			Name:       "vault.enable-graph",
			InnerField: "enableGraph",
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
			Name:     "id",
			Required: true,
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
			Name:     "id",
			Required: true,
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
		&requestflag.Flag[any]{
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
	Usage:   "List matters for the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
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

func handleMattersV1Create(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1NewParams{}

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

	params := githubcomcasemarkcasedevgo.MatterV1UpdateParams{}

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

	params := githubcomcasemarkcasedevgo.MatterV1ListParams{}

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

	return client.Matters.V1.List(ctx, params, options...)
}
