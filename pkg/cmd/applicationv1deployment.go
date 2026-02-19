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
	Usage:   "Trigger a new deployment for a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID",
			Required: true,
			BodyPath: "projectId",
		},
		&requestflag.Flag[string]{
			Name:     "ref",
			Usage:    "Git ref (branch, tag, or commit) to deploy",
			BodyPath: "ref",
		},
		&requestflag.Flag[string]{
			Name:     "target",
			Usage:    "Deployment target",
			Default:  "production",
			BodyPath: "target",
		},
	},
	Action:          handleApplicationsV1DeploymentsCreate,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific deployment including build logs",
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
		&requestflag.Flag[bool]{
			Name:      "include-logs",
			Usage:     "Include build logs",
			Default:   false,
			QueryPath: "includeLogs",
		},
	},
	Action:          handleApplicationsV1DeploymentsRetrieve,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsList = cli.Command{
	Name:    "list",
	Usage:   "List deployments for a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID",
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
			Usage:     "Filter by deployment state",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "target",
			Usage:     "Filter by deployment target",
			QueryPath: "target",
		},
	},
	Action:          handleApplicationsV1DeploymentsList,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a running deployment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID (for authorization)",
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
	Usage:   "Get build logs for a specific deployment",
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
	},
	Action:          handleApplicationsV1DeploymentsGetLogs,
	HideHelpCommand: true,
}

var applicationsV1DeploymentsGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Get the current status of a deployment",
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
