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

var mattersV1SharesCreate = cli.Command{
	Name:    "create",
	Usage:   "Grant another organization scoped access to this matter and its primary vault.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "target-org-id",
			Required: true,
			BodyPath: "target_org_id",
		},
		&requestflag.Flag[any]{
			Name:     "expires-at",
			BodyPath: "expires_at",
		},
		&requestflag.Flag[string]{
			Name:     "permission",
			Usage:    `Allowed values: "read", "edit".`,
			Default:  "read",
			BodyPath: "permission",
		},
	},
	Action:          handleMattersV1SharesCreate,
	HideHelpCommand: true,
}

var mattersV1SharesList = cli.Command{
	Name:    "list",
	Usage:   "List cross-org shares for a matter. Owner only.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMattersV1SharesList,
	HideHelpCommand: true,
}

var mattersV1SharesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Revoke a matter share and its linked vault share.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "share-id",
			Required: true,
		},
	},
	Action:          handleMattersV1SharesDelete,
	HideHelpCommand: true,
}

func handleMattersV1SharesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1ShareNewParams{}

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

	return client.Matters.V1.Shares.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1SharesList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Matters.V1.Shares.List(ctx, cmd.Value("id").(string), options...)
}

func handleMattersV1SharesDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("share-id") && len(unusedArgs) > 0 {
		cmd.Set("share-id", unusedArgs[0])
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

	return client.Matters.V1.Shares.Delete(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("share-id").(string),
		options...,
	)
}
