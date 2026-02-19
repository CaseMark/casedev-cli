// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var applicationsV1ProjectsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new web application project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "git-repo",
			Usage:    `GitHub repository URL or "owner/repo"`,
			Required: true,
			BodyPath: "gitRepo",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Project name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "build-command",
			Usage:    "Custom build command",
			BodyPath: "buildCommand",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "environment-variable",
			Usage:    "Environment variables to set on the project",
			BodyPath: "environmentVariables",
		},
		&requestflag.Flag[string]{
			Name:     "framework",
			Usage:    `Framework (e.g., "nextjs", "remix", "astro")`,
			BodyPath: "framework",
		},
		&requestflag.Flag[string]{
			Name:     "git-branch",
			Usage:    "Git branch to deploy",
			Default:  "main",
			BodyPath: "gitBranch",
		},
		&requestflag.Flag[string]{
			Name:     "install-command",
			Usage:    "Custom install command",
			BodyPath: "installCommand",
		},
		&requestflag.Flag[string]{
			Name:     "output-directory",
			Usage:    "Build output directory",
			BodyPath: "outputDirectory",
		},
		&requestflag.Flag[string]{
			Name:     "root-directory",
			Usage:    "Root directory of the project",
			BodyPath: "rootDirectory",
		},
	},
	Action:          handleApplicationsV1ProjectsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"environment-variable": {
		&requestflag.InnerFlag[string]{
			Name:       "environment-variable.key",
			Usage:      "Environment variable name",
			InnerField: "key",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "environment-variable.target",
			Usage:      "Deployment targets for this variable",
			InnerField: "target",
		},
		&requestflag.InnerFlag[string]{
			Name:       "environment-variable.value",
			Usage:      "Environment variable value",
			InnerField: "value",
		},
		&requestflag.InnerFlag[string]{
			Name:       "environment-variable.type",
			Usage:      "Variable type",
			InnerField: "type",
		},
	},
})

var applicationsV1ProjectsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific web application project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleApplicationsV1ProjectsRetrieve,
	HideHelpCommand: true,
}

var applicationsV1ProjectsList = cli.Command{
	Name:            "list",
	Usage:           "List all web application projects",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleApplicationsV1ProjectsList,
	HideHelpCommand: true,
}

var applicationsV1ProjectsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a web application project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "delete-from-hosting",
			Usage:     "Also delete the project from hosting (default: true)",
			Default:   true,
			QueryPath: "deleteFromHosting",
		},
	},
	Action:          handleApplicationsV1ProjectsDelete,
	HideHelpCommand: true,
}

var applicationsV1ProjectsCreateDeployment = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-deployment",
	Usage:   "Trigger a new deployment for a project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "environment-variable",
			Usage:    "Additional environment variables to set or update before deployment",
			BodyPath: "environmentVariables",
		},
	},
	Action:          handleApplicationsV1ProjectsCreateDeployment,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"environment-variable": {
		&requestflag.InnerFlag[string]{
			Name:       "environment-variable.key",
			Usage:      "Environment variable name",
			InnerField: "key",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "environment-variable.target",
			Usage:      "Deployment targets for this variable",
			InnerField: "target",
		},
		&requestflag.InnerFlag[string]{
			Name:       "environment-variable.value",
			Usage:      "Environment variable value",
			InnerField: "value",
		},
		&requestflag.InnerFlag[string]{
			Name:       "environment-variable.type",
			Usage:      "Variable type",
			InnerField: "type",
		},
	},
})

var applicationsV1ProjectsCreateDomain = cli.Command{
	Name:    "create-domain",
	Usage:   "Add a custom domain to a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    "Domain name to add",
			Required: true,
			BodyPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:     "git-branch",
			Usage:    "Git branch to associate with this domain",
			BodyPath: "gitBranch",
		},
	},
	Action:          handleApplicationsV1ProjectsCreateDomain,
	HideHelpCommand: true,
}

var applicationsV1ProjectsCreateEnv = cli.Command{
	Name:    "create-env",
	Usage:   "Create a new environment variable for a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "key",
			Usage:    "Environment variable name",
			Required: true,
			BodyPath: "key",
		},
		&requestflag.Flag[[]string]{
			Name:     "target",
			Usage:    "Deployment targets for this variable",
			Required: true,
			BodyPath: "target",
		},
		&requestflag.Flag[string]{
			Name:     "value",
			Usage:    "Environment variable value",
			Required: true,
			BodyPath: "value",
		},
		&requestflag.Flag[string]{
			Name:     "git-branch",
			Usage:    "Specific git branch (for preview deployments)",
			BodyPath: "gitBranch",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Variable type",
			Default:  "encrypted",
			BodyPath: "type",
		},
	},
	Action:          handleApplicationsV1ProjectsCreateEnv,
	HideHelpCommand: true,
}

var applicationsV1ProjectsDeleteDomain = cli.Command{
	Name:    "delete-domain",
	Usage:   "Remove a domain from a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "domain",
			Required: true,
		},
	},
	Action:          handleApplicationsV1ProjectsDeleteDomain,
	HideHelpCommand: true,
}

var applicationsV1ProjectsDeleteEnv = cli.Command{
	Name:    "delete-env",
	Usage:   "Delete an environment variable from a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "env-id",
			Required: true,
		},
	},
	Action:          handleApplicationsV1ProjectsDeleteEnv,
	HideHelpCommand: true,
}

var applicationsV1ProjectsGetRuntimeLogs = cli.Command{
	Name:    "get-runtime-logs",
	Usage:   "Get runtime/function logs for a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[float64]{
			Name:      "limit",
			Usage:     "Maximum number of logs to return",
			Default:   100,
			QueryPath: "limit",
		},
	},
	Action:          handleApplicationsV1ProjectsGetRuntimeLogs,
	HideHelpCommand: true,
}

var applicationsV1ProjectsListDeployments = cli.Command{
	Name:    "list-deployments",
	Usage:   "List deployments for a specific project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
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
	Action:          handleApplicationsV1ProjectsListDeployments,
	HideHelpCommand: true,
}

var applicationsV1ProjectsListDomains = cli.Command{
	Name:    "list-domains",
	Usage:   "List all domains configured for a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleApplicationsV1ProjectsListDomains,
	HideHelpCommand: true,
}

var applicationsV1ProjectsListEnv = cli.Command{
	Name:    "list-env",
	Usage:   "List all environment variables for a project (values are hidden unless\ndecrypt=true)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "decrypt",
			Usage:     "Whether to decrypt and return values (requires appropriate permissions)",
			Default:   false,
			QueryPath: "decrypt",
		},
	},
	Action:          handleApplicationsV1ProjectsListEnv,
	HideHelpCommand: true,
}

func handleApplicationsV1ProjectsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectNewParams{}

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

	return client.Applications.V1.Projects.New(ctx, params, options...)
}

func handleApplicationsV1ProjectsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Applications.V1.Projects.Get(ctx, cmd.Value("id").(string), options...)
}

func handleApplicationsV1ProjectsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Applications.V1.Projects.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "applications:v1:projects list", obj, format, transform)
}

func handleApplicationsV1ProjectsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectDeleteParams{}

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

	return client.Applications.V1.Projects.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1ProjectsCreateDeployment(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDeploymentParams{}

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

	return client.Applications.V1.Projects.NewDeployment(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1ProjectsCreateDomain(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDomainParams{}

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

	return client.Applications.V1.Projects.NewDomain(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1ProjectsCreateEnv(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectNewEnvParams{}

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

	return client.Applications.V1.Projects.NewEnv(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1ProjectsDeleteDomain(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("domain") && len(unusedArgs) > 0 {
		cmd.Set("domain", unusedArgs[0])
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

	return client.Applications.V1.Projects.DeleteDomain(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("domain").(string),
		options...,
	)
}

func handleApplicationsV1ProjectsDeleteEnv(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("env-id") && len(unusedArgs) > 0 {
		cmd.Set("env-id", unusedArgs[0])
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

	return client.Applications.V1.Projects.DeleteEnv(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("env-id").(string),
		options...,
	)
}

func handleApplicationsV1ProjectsGetRuntimeLogs(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectGetRuntimeLogsParams{}

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

	return client.Applications.V1.Projects.GetRuntimeLogs(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1ProjectsListDeployments(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectListDeploymentsParams{}

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

	return client.Applications.V1.Projects.ListDeployments(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleApplicationsV1ProjectsListDomains(ctx context.Context, cmd *cli.Command) error {
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

	return client.Applications.V1.Projects.ListDomains(ctx, cmd.Value("id").(string), options...)
}

func handleApplicationsV1ProjectsListEnv(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ApplicationV1ProjectListEnvParams{}

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

	return client.Applications.V1.Projects.ListEnv(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
