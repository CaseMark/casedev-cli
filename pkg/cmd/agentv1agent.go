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

var agentV1AgentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new agent definition with a scoped API key. The agent can then be used\nto create and execute runs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "System instructions that define agent behavior",
			Required: true,
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Display name for the agent",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description of the agent",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "disabled-tool",
			Usage:    "Denylist of tools the agent cannot use. Mutually exclusive with enabledTools — set one or the other, not both.",
			BodyPath: "disabledTools",
		},
		&requestflag.Flag[any]{
			Name:     "enabled-tool",
			Usage:    "Allowlist of tools the agent can use. Mutually exclusive with disabledTools — set one or the other, not both.",
			BodyPath: "enabledTools",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "LLM model identifier (e.g. anthropic/claude-sonnet-4.6). Defaults to anthropic/claude-sonnet-4.6",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "sandbox",
			Usage:    "Custom sandbox configuration (cpu, memoryMiB)",
			BodyPath: "sandbox",
		},
		&requestflag.Flag[any]{
			Name:     "vault-group",
			Usage:    "Restrict agent to vaults within specific vault group IDs",
			BodyPath: "vaultGroups",
		},
		&requestflag.Flag[any]{
			Name:     "vault-id",
			Usage:    "Restrict agent to specific vault IDs",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleAgentV1AgentsCreate,
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

var agentV1AgentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single agent definition by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1AgentsRetrieve,
	HideHelpCommand: true,
}

var agentV1AgentsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an agent definition. Only provided fields are changed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Updated agent description. Pass null to clear if supported by the client.",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "disabled-tool",
			Usage:    "Denylist of tools the agent cannot use. Mutually exclusive with enabledTools — set one or the other, not both. Pass null to clear.",
			BodyPath: "disabledTools",
		},
		&requestflag.Flag[any]{
			Name:     "enabled-tool",
			Usage:    "Allowlist of tools the agent can use. Mutually exclusive with disabledTools — set one or the other, not both. Pass null to clear.",
			BodyPath: "enabledTools",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Updated system instructions that guide agent behavior",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model identifier the agent should use for future runs",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Updated agent display name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sandbox",
			Usage:    "Sandbox configuration override for future agent runs. Pass null to clear.",
			BodyPath: "sandbox",
		},
		&requestflag.Flag[any]{
			Name:     "vault-group",
			Usage:    "Vault group IDs the agent can access. Pass null to clear.",
			BodyPath: "vaultGroups",
		},
		&requestflag.Flag[any]{
			Name:     "vault-id",
			Usage:    "Vault IDs the agent can access directly. Pass null to clear.",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleAgentV1AgentsUpdate,
	HideHelpCommand: true,
}

var agentV1AgentsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all active agents for the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor (agent ID from previous page). Returns agents created before this agent.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of agents to return (default 50, max 250)",
			Default:   50,
			QueryPath: "limit",
		},
	},
	Action:          handleAgentV1AgentsList,
	HideHelpCommand: true,
}

var agentV1AgentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-deletes an agent and revokes its scoped API key.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1AgentsDelete,
	HideHelpCommand: true,
}

func handleAgentV1AgentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1AgentNewParams{}

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
	_, err = client.Agent.V1.Agents.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:agents create", obj, format, transform)
}

func handleAgentV1AgentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Agents.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:agents retrieve", obj, format, transform)
}

func handleAgentV1AgentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1AgentUpdateParams{}

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
	_, err = client.Agent.V1.Agents.Update(
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
	return ShowJSON(os.Stdout, "agent:v1:agents update", obj, format, transform)
}

func handleAgentV1AgentsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1AgentListParams{}

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
	_, err = client.Agent.V1.Agents.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:agents list", obj, format, transform)
}

func handleAgentV1AgentsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Agents.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:agents delete", obj, format, transform)
}
