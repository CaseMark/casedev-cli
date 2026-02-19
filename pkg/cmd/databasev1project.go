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

var databaseV1ProjectsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new serverless Postgres database project powered by Neon. Includes\nautomatic scaling, connection pooling, and a default 'main' branch with 'neondb'\ndatabase. Supports branching for isolated dev/staging environments. Perfect for\ncase management applications, document workflows, and litigation support\nsystems.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Project name (letters, numbers, hyphens, underscores only)",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional project description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "AWS region for database deployment",
			Default:  "aws-us-east-1",
			BodyPath: "region",
		},
	},
	Action:          handleDatabaseV1ProjectsCreate,
	HideHelpCommand: true,
}

var databaseV1ProjectsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves detailed information about a specific database project including\nbranches, databases, storage/compute metrics, connection host, and linked\ndeployments. Fetches live usage statistics from Neon API.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDatabaseV1ProjectsRetrieve,
	HideHelpCommand: true,
}

var databaseV1ProjectsList = cli.Command{
	Name:            "list",
	Usage:           "Retrieves all serverless Postgres database projects for the authenticated\norganization. Includes storage and compute metrics, plus linked deployments from\nThurgood apps and Compute instances.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDatabaseV1ProjectsList,
	HideHelpCommand: true,
}

var databaseV1ProjectsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a database project from Neon and marks it as deleted in\nCase.dev. This action cannot be undone and will destroy all data including\nbranches and databases. Use with caution.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDatabaseV1ProjectsDelete,
	HideHelpCommand: true,
}

var databaseV1ProjectsCreateBranch = cli.Command{
	Name:    "create-branch",
	Usage:   "Creates a new branch from the specified parent branch (or default 'main'\nbranch). Branches provide instant point-in-time clones of your database for\nisolated development, staging, testing, or feature work. Perfect for testing\nschema changes, running migrations safely, or creating ephemeral preview\nenvironments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Branch name (letters, numbers, hyphens, underscores only)",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "parent-branch-id",
			Usage:    "Parent branch ID to clone from (defaults to main branch)",
			BodyPath: "parentBranchId",
		},
	},
	Action:          handleDatabaseV1ProjectsCreateBranch,
	HideHelpCommand: true,
}

var databaseV1ProjectsGetConnection = cli.Command{
	Name:    "get-connection",
	Usage:   "Retrieves the PostgreSQL connection URI for a database project. Supports\nselecting specific branches and pooled vs direct connections. Connection strings\ninclude credentials and should be stored securely. Use for configuring\napplications and deployment environments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "branch",
			Usage:     "Branch name (defaults to 'main')",
			QueryPath: "branch",
		},
		&requestflag.Flag[bool]{
			Name:      "pooled",
			Usage:     "Use pooled connection (PgBouncer)",
			QueryPath: "pooled",
		},
	},
	Action:          handleDatabaseV1ProjectsGetConnection,
	HideHelpCommand: true,
}

var databaseV1ProjectsListBranches = cli.Command{
	Name:    "list-branches",
	Usage:   "Retrieves all branches for a database project. Branches enable isolated\ndevelopment and testing environments with instant point-in-time cloning. Each\nbranch includes the default branch and any custom branches created for staging,\ntesting, or feature development.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDatabaseV1ProjectsListBranches,
	HideHelpCommand: true,
}

func handleDatabaseV1ProjectsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.DatabaseV1ProjectNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Database.V1.Projects.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects create", obj, format, transform)
}

func handleDatabaseV1ProjectsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Database.V1.Projects.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects retrieve", obj, format, transform)
}

func handleDatabaseV1ProjectsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Database.V1.Projects.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects list", obj, format, transform)
}

func handleDatabaseV1ProjectsDelete(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Database.V1.Projects.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects delete", obj, format, transform)
}

func handleDatabaseV1ProjectsCreateBranch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.DatabaseV1ProjectNewBranchParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Database.V1.Projects.NewBranch(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects create-branch", obj, format, transform)
}

func handleDatabaseV1ProjectsGetConnection(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.DatabaseV1ProjectGetConnectionParams{}

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
	_, err = client.Database.V1.Projects.GetConnection(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects get-connection", obj, format, transform)
}

func handleDatabaseV1ProjectsListBranches(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Database.V1.Projects.ListBranches(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "database:v1:projects list-branches", obj, format, transform)
}
