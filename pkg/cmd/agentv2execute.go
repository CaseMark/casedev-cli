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

var agentV2ExecuteCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates an ephemeral agent and executes it immediately. By default this uses the\nlightweight synchronous linc runtime on Vercel Sandbox. Set `agentRuntime: true`\nto opt into the legacy Daytona-backed agent runtime.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[any]{
			Name:     "agent-runtime",
			Usage:    "Set to true to opt into the legacy Daytona-backed agent runtime.",
			BodyPath: "agentRuntime",
		},
		&requestflag.Flag[any]{
			Name:     "disabled-tool",
			BodyPath: "disabledTools",
		},
		&requestflag.Flag[any]{
			Name:     "enabled-tool",
			BodyPath: "enabledTools",
		},
		&requestflag.Flag[any]{
			Name:     "guidance",
			BodyPath: "guidance",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "object-id",
			BodyPath: "objectIds",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sandbox",
			BodyPath: "sandbox",
		},
		&requestflag.Flag[any]{
			Name:     "vault-id",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleAgentV2ExecuteCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sandbox": {
		&requestflag.InnerFlag[int64]{
			Name:       "sandbox.cpu",
			InnerField: "cpu",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "sandbox.memory-mi-b",
			InnerField: "memoryMiB",
		},
	},
})

func handleAgentV2ExecuteCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV2ExecuteNewParams{}

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
	_, err = client.Agent.V2.Execute.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v2:execute create", obj, format, transform)
}
