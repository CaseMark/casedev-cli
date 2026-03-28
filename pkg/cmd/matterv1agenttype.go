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

var mattersV1AgentTypesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a reusable agent role for legal matter orchestration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instructions",
			Required: true,
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]string]{
			Name:     "disabled-tool",
			BodyPath: "disabled_tools",
		},
		&requestflag.Flag[[]string]{
			Name:     "enabled-tool",
			BodyPath: "enabled_tools",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "is_active",
		},
		&requestflag.Flag[bool]{
			Name:     "is-default",
			BodyPath: "is_default",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
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
	Action:          handleMattersV1AgentTypesCreate,
	HideHelpCommand: true,
}

var mattersV1AgentTypesList = cli.Command{
	Name:    "list",
	Usage:   "List reusable agent roles for the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "active",
			QueryPath: "active",
		},
	},
	Action:          handleMattersV1AgentTypesList,
	HideHelpCommand: true,
}

func handleMattersV1AgentTypesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1AgentTypeNewParams{}

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

	return client.Matters.V1.AgentTypes.New(ctx, params, options...)
}

func handleMattersV1AgentTypesList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1AgentTypeListParams{}

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

	return client.Matters.V1.AgentTypes.List(ctx, params, options...)
}
