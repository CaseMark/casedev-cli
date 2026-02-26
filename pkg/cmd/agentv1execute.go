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

var agentV1ExecuteCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates an ephemeral agent and immediately executes a run. Returns the run ID\nfor polling status and results. This is the fastest way to run an agent without\nmanaging agent lifecycle.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Task prompt for the agent",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[any]{
			Name:     "disabled-tool",
			Usage:    "Denylist of tools the agent cannot use",
			BodyPath: "disabledTools",
		},
		&requestflag.Flag[any]{
			Name:     "enabled-tool",
			Usage:    "Allowlist of tools the agent can use",
			BodyPath: "enabledTools",
		},
		&requestflag.Flag[any]{
			Name:     "guidance",
			Usage:    "Additional context or constraints for this run",
			BodyPath: "guidance",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "System instructions. Defaults to a general-purpose legal assistant prompt if not provided.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "LLM model identifier. Defaults to anthropic/claude-sonnet-4.6",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "object-id",
			Usage:    "Scope this run to specific vault object IDs. The agent will only access these objects.",
			BodyPath: "objectIds",
		},
		&requestflag.Flag[any]{
			Name:     "sandbox",
			Usage:    "Custom sandbox resources (cpu, memoryMiB)",
			BodyPath: "sandbox",
		},
		&requestflag.Flag[any]{
			Name:     "vault-id",
			Usage:    "Restrict agent to specific vault IDs",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleAgentV1ExecuteCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sandbox": {
		&requestflag.InnerFlag[int64]{
			Name:       "sandbox.cpu",
			Usage:      "Number of CPUs",
			InnerField: "cpu",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "sandbox.memory-mi-b",
			Usage:      "Memory in MiB",
			InnerField: "memoryMiB",
		},
	},
})

func handleAgentV1ExecuteCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1ExecuteNewParams{}

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
	_, err = client.Agent.V1.Execute.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:execute create", obj, format, transform)
}
