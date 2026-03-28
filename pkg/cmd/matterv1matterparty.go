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

var mattersV1MatterPartiesCreate = cli.Command{
	Name:    "create",
	Usage:   "Attach a reusable party to a matter with a matter-specific role.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "party-id",
			Required: true,
			BodyPath: "party_id",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    `Allowed values: "client", "prospect", "opposing_party", "opposing_counsel", "co_counsel", "judge", "expert", "witness", "vendor", "insurer", "other".`,
			Required: true,
			BodyPath: "role",
		},
		&requestflag.Flag[any]{
			Name:     "custom-fields",
			BodyPath: "custom_fields",
		},
		&requestflag.Flag[bool]{
			Name:     "is-primary",
			BodyPath: "is_primary",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "notes",
			BodyPath: "notes",
		},
		&requestflag.Flag[bool]{
			Name:     "set-as-client",
			BodyPath: "set_as_client",
		},
	},
	Action:          handleMattersV1MatterPartiesCreate,
	HideHelpCommand: true,
}

var mattersV1MatterPartiesList = cli.Command{
	Name:    "list",
	Usage:   "List parties attached to a matter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMattersV1MatterPartiesList,
	HideHelpCommand: true,
}

func handleMattersV1MatterPartiesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1MatterPartyNewParams{}

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

	return client.Matters.V1.MatterParties.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1MatterPartiesList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Matters.V1.MatterParties.List(ctx, cmd.Value("id").(string), options...)
}
