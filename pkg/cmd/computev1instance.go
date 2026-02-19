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

var computeV1InstancesCreate = cli.Command{
	Name:    "create",
	Usage:   "Launches a new GPU compute instance with automatic SSH key generation. Supports\nmounting Case.dev Vaults as filesystems and configurable auto-shutdown. Instance\nboots in ~2-5 minutes. Perfect for batch OCR processing, AI model training, and\nintensive document analysis workloads.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instance-type",
			Usage:    "GPU type (e.g., 'gpu_1x_h100_sxm5')",
			Required: true,
			BodyPath: "instanceType",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Instance name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region (e.g., 'us-west-1')",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[any]{
			Name:     "auto-shutdown-minutes",
			Usage:    "Auto-shutdown timer (null = never)",
			BodyPath: "autoShutdownMinutes",
		},
		&requestflag.Flag[[]string]{
			Name:     "vault-id",
			Usage:    "Vault IDs to mount",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleComputeV1InstancesCreate,
	HideHelpCommand: true,
}

var computeV1InstancesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves detailed information about a GPU instance including SSH connection\ndetails, vault mount scripts, real-time cost tracking, and current status. SSH\nprivate key included for secure access.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleComputeV1InstancesRetrieve,
	HideHelpCommand: true,
}

var computeV1InstancesList = cli.Command{
	Name:            "list",
	Usage:           "Retrieves all GPU compute instances for your organization with real-time status\nupdates from Lambda Labs. Includes pricing, runtime metrics, and auto-shutdown\nconfiguration. Perfect for monitoring AI workloads, document processing jobs,\nand cost tracking.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleComputeV1InstancesList,
	HideHelpCommand: true,
}

var computeV1InstancesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Terminates a running GPU instance, calculates final cost, and cleans up SSH\nkeys. This action is permanent and cannot be undone. All data on the instance\nwill be lost.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleComputeV1InstancesDelete,
	HideHelpCommand: true,
}

func handleComputeV1InstancesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1InstanceNewParams{}

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
	_, err = client.Compute.V1.Instances.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:instances create", obj, format, transform)
}

func handleComputeV1InstancesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Compute.V1.Instances.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:instances retrieve", obj, format, transform)
}

func handleComputeV1InstancesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Compute.V1.Instances.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:instances list", obj, format, transform)
}

func handleComputeV1InstancesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Compute.V1.Instances.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:instances delete", obj, format, transform)
}
