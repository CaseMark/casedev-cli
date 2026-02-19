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

var computeV1EnvironmentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new compute environment for running serverless workloads. Each\nenvironment gets its own isolated namespace with a unique domain for hosting\napplications and APIs. The first environment created becomes the default\nenvironment for the organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Environment name (alphanumeric, hyphens, and underscores only)",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleComputeV1EnvironmentsCreate,
	HideHelpCommand: true,
}

var computeV1EnvironmentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific compute environment by name. Returns environment\nconfiguration including status, domain, and metadata for your serverless compute\ninfrastructure.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
		},
	},
	Action:          handleComputeV1EnvironmentsRetrieve,
	HideHelpCommand: true,
}

var computeV1EnvironmentsList = cli.Command{
	Name:            "list",
	Usage:           "Retrieve all compute environments for your organization. Environments provide\nisolated execution contexts for running code and workflows.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleComputeV1EnvironmentsList,
	HideHelpCommand: true,
}

var computeV1EnvironmentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently delete a compute environment and all its associated resources. This\nwill stop all running deployments and clean up related configurations. The\ndefault environment cannot be deleted if other environments exist.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
		},
	},
	Action:          handleComputeV1EnvironmentsDelete,
	HideHelpCommand: true,
}

var computeV1EnvironmentsSetDefault = cli.Command{
	Name:    "set-default",
	Usage:   "Sets a compute environment as the default for the organization. Only one\nenvironment can be default at a time - setting a new default will automatically\nunset the previous one.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
		},
	},
	Action:          handleComputeV1EnvironmentsSetDefault,
	HideHelpCommand: true,
}

func handleComputeV1EnvironmentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1EnvironmentNewParams{}

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
	_, err = client.Compute.V1.Environments.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:environments create", obj, format, transform)
}

func handleComputeV1EnvironmentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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
	_, err = client.Compute.V1.Environments.Get(ctx, cmd.Value("name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:environments retrieve", obj, format, transform)
}

func handleComputeV1EnvironmentsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Compute.V1.Environments.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:environments list", obj, format, transform)
}

func handleComputeV1EnvironmentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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
	_, err = client.Compute.V1.Environments.Delete(ctx, cmd.Value("name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:environments delete", obj, format, transform)
}

func handleComputeV1EnvironmentsSetDefault(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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
	_, err = client.Compute.V1.Environments.SetDefault(ctx, cmd.Value("name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:environments set-default", obj, format, transform)
}
