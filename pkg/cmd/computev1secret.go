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

var computeV1SecretsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new secret group in a compute environment. Secret groups organize\nrelated secrets for use in serverless functions and workflows. If no environment\nis specified, the group is created in the default environment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Unique name for the secret group. Must contain only letters, numbers, hyphens, and underscores.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description of the secret group's purpose",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "env",
			Usage:    "Environment name where the secret group will be created. Uses default environment if not specified.",
			BodyPath: "env",
		},
	},
	Action:          handleComputeV1SecretsCreate,
	HideHelpCommand: true,
}

var computeV1SecretsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve all secret groups for a compute environment. Secret groups organize\nrelated secrets (API keys, credentials, etc.) that can be securely accessed by\ncompute jobs during execution.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "env",
			Usage:     "Environment name to list secret groups for. If not specified, uses the default environment.",
			QueryPath: "env",
		},
	},
	Action:          handleComputeV1SecretsList,
	HideHelpCommand: true,
}

var computeV1SecretsDeleteGroup = cli.Command{
	Name:    "delete-group",
	Usage:   "Delete an entire secret group or a specific key within a secret group. When\ndeleting a specific key, the remaining secrets in the group are preserved. When\ndeleting the entire group, all secrets and the group itself are removed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "env",
			Usage:     "Environment name. If not provided, uses the default environment",
			QueryPath: "env",
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Usage:     "Specific key to delete within the group. If not provided, the entire group is deleted",
			QueryPath: "key",
		},
	},
	Action:          handleComputeV1SecretsDeleteGroup,
	HideHelpCommand: true,
}

var computeV1SecretsRetrieveGroup = cli.Command{
	Name:    "retrieve-group",
	Usage:   "Retrieve the keys (names) of secrets in a specified group within a compute\nenvironment. For security reasons, actual secret values are not returned - only\nthe keys and metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "env",
			Usage:     "Environment name. If not specified, uses the default environment",
			QueryPath: "env",
		},
	},
	Action:          handleComputeV1SecretsRetrieveGroup,
	HideHelpCommand: true,
}

var computeV1SecretsUpdateGroup = cli.Command{
	Name:    "update-group",
	Usage:   "Set or update secrets in a compute secret group. Secrets are encrypted with\nAES-256-GCM. Use this to manage environment variables and API keys for your\ncompute workloads.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "secrets",
			Usage:    "Key-value pairs of secrets to set",
			Required: true,
			BodyPath: "secrets",
		},
		&requestflag.Flag[string]{
			Name:     "env",
			Usage:    "Environment name (optional, uses default if not specified)",
			BodyPath: "env",
		},
	},
	Action:          handleComputeV1SecretsUpdateGroup,
	HideHelpCommand: true,
}

func handleComputeV1SecretsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1SecretNewParams{}

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
	_, err = client.Compute.V1.Secrets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:secrets create", obj, format, transform)
}

func handleComputeV1SecretsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1SecretListParams{}

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
	_, err = client.Compute.V1.Secrets.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:secrets list", obj, format, transform)
}

func handleComputeV1SecretsDeleteGroup(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group") && len(unusedArgs) > 0 {
		cmd.Set("group", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1SecretDeleteGroupParams{}

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
	_, err = client.Compute.V1.Secrets.DeleteGroup(
		ctx,
		cmd.Value("group").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:secrets delete-group", obj, format, transform)
}

func handleComputeV1SecretsRetrieveGroup(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group") && len(unusedArgs) > 0 {
		cmd.Set("group", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1SecretGetGroupParams{}

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
	_, err = client.Compute.V1.Secrets.GetGroup(
		ctx,
		cmd.Value("group").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:secrets retrieve-group", obj, format, transform)
}

func handleComputeV1SecretsUpdateGroup(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group") && len(unusedArgs) > 0 {
		cmd.Set("group", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1SecretUpdateGroupParams{}

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
	_, err = client.Compute.V1.Secrets.UpdateGroup(
		ctx,
		cmd.Value("group").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:secrets update-group", obj, format, transform)
}
