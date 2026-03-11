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

var applicationsV1DeploymentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a deployment for an existing project by fetching repository files from\nGitHub and uploading them to the hosting provider. Use ref to deploy a branch,\ntag, or commit other than the project default branch.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID to deploy",
			Required: true,
			BodyPath: "projectId",
		},
		&requestflag.Flag[string]{
			Name:     "ref",
			Usage:    "Git branch, tag, or commit to deploy. Defaults to the project branch.",
			BodyPath: "ref",
		},
		&requestflag.Flag[string]{
			Name:     "target",
			Usage:    "Deployment target environment",
			Default:  "production",
			BodyPath: "target",
		},
	},
	Action:          handleApplicationsV1DeploymentsCreate,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns deployment details for one project in the authenticated organization.\nSet includeLogs=true to include recent build output in the response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID used to verify access to the deployment",
			Required:  true,
			QueryPath: "projectId",
		},
		&requestflag.Flag[bool]{
			Name:      "include-logs",
			Usage:     "Whether to include build logs in the response",
			Default:   false,
			QueryPath: "includeLogs",
		},
	},
	Action:          handleApplicationsV1DeploymentsRetrieve,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsList = cli.Command{
	Name:    "list",
	Usage:   "Lists recent deployments for one project in the authenticated organization. Use\nthe optional filters to narrow results by target or deployment state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID to list deployments for",
			Required:  true,
			QueryPath: "projectId",
		},
		&requestflag.Flag[float64]{
			Name:      "limit",
			Usage:     "Maximum number of deployments to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Deployment state to filter by",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "target",
			Usage:     "Deployment target to filter by",
			QueryPath: "target",
		},
	},
	Action:          handleApplicationsV1DeploymentsList,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels a running deployment after verifying that the referenced project belongs\nto the authenticated organization. Use this when a build is stuck,\nmisconfigured, or no longer needed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID used to verify access to the deployment",
			Required: true,
			BodyPath: "projectId",
		},
	},
	Action:          handleApplicationsV1DeploymentsCancel,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsCreateFromFiles = cli.Command{
	Name:            "create-from-files",
	Usage:           "Create a deployment from raw file contents (for Thurgood sandbox deployments)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleApplicationsV1DeploymentsCreateFromFiles,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsGetLogs = cli.Command{
	Name:    "get-logs",
	Usage:   "Returns build and runtime log events for a deployment after verifying access to\nthe owning project. Use this when you need detailed output for a failed or\nin-progress build.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID used to verify access to the deployment",
			Required:  true,
			QueryPath: "projectId",
		},
	},
	Action:          handleApplicationsV1DeploymentsGetLogs,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Returns the current status of a deployment without fetching full build logs. Use\nthis endpoint for lightweight polling while a deployment is building or waiting\nto become ready.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleApplicationsV1DeploymentsGetStatus,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsStream = cli.Command{
	Name:    "stream",
	Usage:   "Stream real-time deployment progress events via Server-Sent Events",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID (for authorization)",
			Required:  true,
			QueryPath: "projectId",
		},
		&requestflag.Flag[float64]{
			Name:      "start-index",
			Usage:     "Resume stream from this index (for reconnection)",
			QueryPath: "startIndex",
		},
	},
	Action:          handleApplicationsV1DeploymentsStream,
	HideHelpCommand: true,
}

func handleApplicationsV1DeploymentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1DeploymentNewParams{}

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

	return client.Applications.V1.Deployments.New(ctx, params, options...)
}

func handleApplicationsV1DeploymentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1DeploymentGetParams{}

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

	return client.Applications.V1.Deployments.Get(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1DeploymentsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1DeploymentListParams{}

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

	return client.Applications.V1.Deployments.List(ctx, params, options...)
}

func handleApplicationsV1DeploymentsCancel(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1DeploymentCancelParams{}

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

	return client.Applications.V1.Deployments.Cancel(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1DeploymentsCreateFromFiles(ctx context.Context, cmd *cli.Command) error {
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

	return client.Applications.V1.Deployments.NewFromFiles(ctx, options...)
}

func handleApplicationsV1DeploymentsGetLogs(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1DeploymentGetLogsParams{}

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

	return client.Applications.V1.Deployments.GetLogs(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1DeploymentsGetStatus(ctx context.Context, cmd *cli.Command) error {
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

	return client.Applications.V1.Deployments.GetStatus(ctx, cmd.Value("id").(string), options...)
}

func handleApplicationsV1DeploymentsStream(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1DeploymentStreamParams{}

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

	return client.Applications.V1.Deployments.Stream(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
