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

var mattersV1WorkItemsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an active work item on a matter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[string]{
			Name:     "assignee-id",
			BodyPath: "assignee_id",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "due-at",
			BodyPath: "due_at",
		},
		&requestflag.Flag[[]string]{
			Name:     "exit-criterion",
			BodyPath: "exit_criteria",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "priority",
			Usage:    `Allowed values: "low", "normal", "high", "urgent".`,
			BodyPath: "priority",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "task", "deadline", "review", "filing", "communication", "research", "drafting", "collection", "intake".`,
			BodyPath: "type",
		},
	},
	Action:          handleMattersV1WorkItemsCreate,
	HideHelpCommand: true,
}

var mattersV1WorkItemsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single work item for a matter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "work-item-id",
			Required: true,
		},
	},
	Action:          handleMattersV1WorkItemsRetrieve,
	HideHelpCommand: true,
}

var mattersV1WorkItemsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a matter work item.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "work-item-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "assignee-id",
			BodyPath: "assignee_id",
		},
		&requestflag.Flag[any]{
			Name:     "completed-at",
			BodyPath: "completed_at",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "due-at",
			BodyPath: "due_at",
		},
		&requestflag.Flag[[]string]{
			Name:     "exit-criterion",
			BodyPath: "exit_criteria",
		},
		&requestflag.Flag[any]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "priority",
			Usage:    `Allowed values: "low", "normal", "high", "urgent".`,
			BodyPath: "priority",
		},
		&requestflag.Flag[any]{
			Name:     "started-at",
			BodyPath: "started_at",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "draft", "queued", "in_progress", "blocked", "in_review", "awaiting_human", "done", "canceled".`,
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			BodyPath: "title",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "task", "deadline", "review", "filing", "communication", "research", "drafting", "collection", "intake".`,
			BodyPath: "type",
		},
	},
	Action:          handleMattersV1WorkItemsUpdate,
	HideHelpCommand: true,
}

var mattersV1WorkItemsList = cli.Command{
	Name:    "list",
	Usage:   "List active work items for a matter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "assignee-id",
			QueryPath: "assignee_id",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			QueryPath: "status",
		},
	},
	Action:          handleMattersV1WorkItemsList,
	HideHelpCommand: true,
}

var mattersV1WorkItemsDecide = cli.Command{
	Name:    "decide",
	Usage:   "Approve, revise, block, or reassign a work item. Used by humans or agents to\nmove work items through their lifecycle.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "work-item-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "decision",
			Usage:    `Allowed values: "approve", "revise", "block", "reassign".`,
			Required: true,
			BodyPath: "decision",
		},
		&requestflag.Flag[any]{
			Name:     "agent-type-id",
			BodyPath: "agent_type_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "reason",
			BodyPath: "reason",
		},
	},
	Action:          handleMattersV1WorkItemsDecide,
	HideHelpCommand: true,
}

var mattersV1WorkItemsListExecutions = cli.Command{
	Name:    "list-executions",
	Usage:   "List execution attempts for a work item, including agent and run linkage.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "work-item-id",
			Required: true,
		},
	},
	Action:          handleMattersV1WorkItemsListExecutions,
	HideHelpCommand: true,
}

func handleMattersV1WorkItemsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1WorkItemNewParams{}

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

	return client.Matters.V1.WorkItems.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1WorkItemsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("work-item-id") && len(unusedArgs) > 0 {
		cmd.Set("work-item-id", unusedArgs[0])
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

	return client.Matters.V1.WorkItems.Get(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("work-item-id").(string),
		options...,
	)
}

func handleMattersV1WorkItemsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("work-item-id") && len(unusedArgs) > 0 {
		cmd.Set("work-item-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1WorkItemUpdateParams{}

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

	return client.Matters.V1.WorkItems.Update(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("work-item-id").(string),
		params,
		options...,
	)
}

func handleMattersV1WorkItemsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1WorkItemListParams{}

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

	return client.Matters.V1.WorkItems.List(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleMattersV1WorkItemsDecide(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("work-item-id") && len(unusedArgs) > 0 {
		cmd.Set("work-item-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MatterV1WorkItemDecideParams{}

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

	return client.Matters.V1.WorkItems.Decide(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("work-item-id").(string),
		params,
		options...,
	)
}

func handleMattersV1WorkItemsListExecutions(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("work-item-id") && len(unusedArgs) > 0 {
		cmd.Set("work-item-id", unusedArgs[0])
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

	return client.Matters.V1.WorkItems.ListExecutions(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("work-item-id").(string),
		options...,
	)
}
