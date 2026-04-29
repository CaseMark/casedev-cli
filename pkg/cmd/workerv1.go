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

var workerV1Create = cli.Command{
	Name:            "create",
	Usage:           "Creates a Daytona-backed worker runtime. The worker exposes its native runtime\nAPI through /worker/v1/:id/\\* without reshaping payloads or events.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWorkerV1Create,
	HideHelpCommand: true,
}

var workerV1Retrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get worker",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWorkerV1Retrieve,
	HideHelpCommand: true,
}

var workerV1Delete = cli.Command{
	Name:    "delete",
	Usage:   "End worker",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWorkerV1Delete,
	HideHelpCommand: true,
}

var workerV1ProxyDelete = cli.Command{
	Name:    "proxy-delete",
	Usage:   "Forwards a DELETE request to the worker runtime without translating response\nshapes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "worker-path",
			Required: true,
		},
	},
	Action:          handleWorkerV1ProxyDelete,
	HideHelpCommand: true,
}

var workerV1ProxyGet = cli.Command{
	Name:    "proxy-get",
	Usage:   "Forwards a GET request to the worker runtime without translating response or SSE\nevent shapes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "worker-path",
			Required: true,
		},
	},
	Action:          handleWorkerV1ProxyGet,
	HideHelpCommand: true,
}

var workerV1ProxyPatch = cli.Command{
	Name:    "proxy-patch",
	Usage:   "Forwards a PATCH request to the worker runtime without translating request or\nresponse shapes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "worker-path",
			Required: true,
		},
	},
	Action:          handleWorkerV1ProxyPatch,
	HideHelpCommand: true,
}

var workerV1ProxyPost = cli.Command{
	Name:    "proxy-post",
	Usage:   "Forwards a POST request to the worker runtime without translating request,\nresponse, or SSE event shapes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "worker-path",
			Required: true,
		},
	},
	Action:          handleWorkerV1ProxyPost,
	HideHelpCommand: true,
}

var workerV1ProxyPut = cli.Command{
	Name:    "proxy-put",
	Usage:   "Forwards a PUT request to the worker runtime without translating request or\nresponse shapes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "worker-path",
			Required: true,
		},
	},
	Action:          handleWorkerV1ProxyPut,
	HideHelpCommand: true,
}

func handleWorkerV1Create(ctx context.Context, cmd *cli.Command) error {
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

	return client.Worker.V1.New(ctx, options...)
}

func handleWorkerV1Retrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Worker.V1.Get(ctx, cmd.Value("id").(string), options...)
}

func handleWorkerV1Delete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Worker.V1.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleWorkerV1ProxyDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("worker-path") && len(unusedArgs) > 0 {
		cmd.Set("worker-path", unusedArgs[0])
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

	return client.Worker.V1.ProxyDelete(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("worker-path").(string),
		options...,
	)
}

func handleWorkerV1ProxyGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("worker-path") && len(unusedArgs) > 0 {
		cmd.Set("worker-path", unusedArgs[0])
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

	return client.Worker.V1.ProxyGet(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("worker-path").(string),
		options...,
	)
}

func handleWorkerV1ProxyPatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("worker-path") && len(unusedArgs) > 0 {
		cmd.Set("worker-path", unusedArgs[0])
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

	return client.Worker.V1.ProxyPatch(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("worker-path").(string),
		options...,
	)
}

func handleWorkerV1ProxyPost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("worker-path") && len(unusedArgs) > 0 {
		cmd.Set("worker-path", unusedArgs[0])
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

	return client.Worker.V1.ProxyPost(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("worker-path").(string),
		options...,
	)
}

func handleWorkerV1ProxyPut(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("worker-path") && len(unusedArgs) > 0 {
		cmd.Set("worker-path", unusedArgs[0])
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

	return client.Worker.V1.ProxyPut(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("worker-path").(string),
		options...,
	)
}
