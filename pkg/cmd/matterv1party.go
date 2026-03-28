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
		&requestflag.Flag[any]{
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
		&requestflag.Flag[any]{
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
			Name:     "party-id",
			Required: true,
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
			Name:     "party-id",
			Required: true,
		},
	},
	Action:          handleMattersV1PartiesUpdate,
	HideHelpCommand: true,
}

var mattersV1PartiesList = cli.Command{
	Name:    "list",
	Usage:   "List reusable legal parties for the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email",
			QueryPath: "email",
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

	params := githubcomcasemarkcasedevgo.MatterV1PartyNewParams{}

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

	params := githubcomcasemarkcasedevgo.MatterV1PartyListParams{}

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

	return client.Matters.V1.Parties.List(ctx, params, options...)
}
