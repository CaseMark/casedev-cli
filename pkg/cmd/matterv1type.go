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

var mattersV1TypesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a matter type with plain-English operating instructions and seeded work.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "default-agent-type-id",
			BodyPath: "default_agent_type_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-metadata",
			BodyPath: "default_metadata",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "default-work-item",
			BodyPath: "default_work_items",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]string]{
			Name:     "exit-criterion",
			BodyPath: "exit_criteria",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[[]string]{
			Name:     "intake-requirement",
			BodyPath: "intake_requirements",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "is_active",
		},
		&requestflag.Flag[string]{
			Name:     "orchestration-mode",
			Usage:    `Allowed values: "auto", "human".`,
			BodyPath: "orchestration_mode",
		},
		&requestflag.Flag[string]{
			Name:     "review-agent-type-id",
			BodyPath: "review_agent_type_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "review-criterion",
			BodyPath: "review_criteria",
		},
		&requestflag.Flag[[]string]{
			Name:     "skill-ref",
			BodyPath: "skill_refs",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			BodyPath: "slug",
		},
	},
	Action:          handleMattersV1TypesCreate,
	HideHelpCommand: true,
}

var mattersV1TypesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single matter type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMattersV1TypesRetrieve,
	HideHelpCommand: true,
}

var mattersV1TypesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a matter type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "default-agent-type-id",
			BodyPath: "default_agent_type_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-metadata",
			BodyPath: "default_metadata",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "default-work-item",
			BodyPath: "default_work_items",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]string]{
			Name:     "exit-criterion",
			BodyPath: "exit_criteria",
		},
		&requestflag.Flag[any]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[[]string]{
			Name:     "intake-requirement",
			BodyPath: "intake_requirements",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "is_active",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "orchestration-mode",
			Usage:    `Allowed values: "auto", "human".`,
			BodyPath: "orchestration_mode",
		},
		&requestflag.Flag[any]{
			Name:     "review-agent-type-id",
			BodyPath: "review_agent_type_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "review-criterion",
			BodyPath: "review_criteria",
		},
		&requestflag.Flag[[]string]{
			Name:     "skill-ref",
			BodyPath: "skill_refs",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			BodyPath: "slug",
		},
	},
	Action:          handleMattersV1TypesUpdate,
	HideHelpCommand: true,
}

var mattersV1TypesList = cli.Command{
	Name:    "list",
	Usage:   "List matter types for the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "active",
			QueryPath: "active",
		},
	},
	Action:          handleMattersV1TypesList,
	HideHelpCommand: true,
}

func handleMattersV1TypesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1TypeNewParams{}

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

	return client.Matters.V1.Types.New(ctx, params, options...)
}

func handleMattersV1TypesRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Matters.V1.Types.Get(ctx, cmd.Value("id").(string), options...)
}

func handleMattersV1TypesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1TypeUpdateParams{}

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

	return client.Matters.V1.Types.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1TypesList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1TypeListParams{}

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

	return client.Matters.V1.Types.List(ctx, params, options...)
}
