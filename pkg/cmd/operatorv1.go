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

var operatorV1Create = cli.Command{
	Name:    "create",
	Usage:   "Provision a new operator instance for the organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Operator name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model to use",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "size",
			Usage:    "Instance size",
			BodyPath: "size",
		},
	},
	Action:          handleOperatorV1Create,
	HideHelpCommand: true,
}

var operatorV1CreateChatCompletion = cli.Command{
	Name:            "create-chat-completion",
	Usage:           "Proxy a chat completion request to the organization's operator instance.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOperatorV1CreateChatCompletion,
	HideHelpCommand: true,
}

var operatorV1CreateResponse = cli.Command{
	Name:            "create-response",
	Usage:           "Proxy a response request to the organization's operator instance.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOperatorV1CreateResponse,
	HideHelpCommand: true,
}

var operatorV1GetStatus = cli.Command{
	Name:            "get-status",
	Usage:           "Get the status of the organization's operator instance.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOperatorV1GetStatus,
	HideHelpCommand: true,
}

func handleOperatorV1Create(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.OperatorV1NewParams{}

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

	return client.Operator.V1.New(ctx, params, options...)
}

func handleOperatorV1CreateChatCompletion(ctx context.Context, cmd *cli.Command) error {
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

	return client.Operator.V1.NewChatCompletion(ctx, options...)
}

func handleOperatorV1CreateResponse(ctx context.Context, cmd *cli.Command) error {
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

	return client.Operator.V1.NewResponse(ctx, options...)
}

func handleOperatorV1GetStatus(ctx context.Context, cmd *cli.Command) error {
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

	return client.Operator.V1.GetStatus(ctx, options...)
}
