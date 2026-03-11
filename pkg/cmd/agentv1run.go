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

var agentV1RunCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a run in queued state. Call POST /agent/v1/run/:id/exec to start\nexecution.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Usage:    "ID of the agent to run",
			Required: true,
			BodyPath: "agentId",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Task prompt for the agent",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[any]{
			Name:     "callback-url",
			Usage:    "HTTPS callback URL to receive a notification when the run completes. Registered atomically with the run — eliminates the race condition of calling /watch after /exec. Additional watchers can still be added via POST /run/:id/watch.",
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[any]{
			Name:     "guidance",
			Usage:    "Additional guidance for this run",
			BodyPath: "guidance",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Override the agent default model for this run",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "object-id",
			Usage:    "Scope this run to specific vault object IDs. The agent will only be able to access these objects during execution.",
			BodyPath: "objectIds",
		},
	},
	Action:          handleAgentV1RunCreate,
	HideHelpCommand: true,
}

var agentV1RunList = cli.Command{
	Name:    "list",
	Usage:   "Lists agent runs for the authenticated organization. Supports filtering by\nagent, status, and cursor-based pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Usage:     "Filter by agent ID",
			QueryPath: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor (run ID from previous page). Returns runs created before this run.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of runs to return (default 50, max 250)",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by run status",
			QueryPath: "status",
		},
	},
	Action:          handleAgentV1RunList,
	HideHelpCommand: true,
}

var agentV1RunCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels a running or queued run. Idempotent — cancelling a finished run returns\nits current status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1RunCancel,
	HideHelpCommand: true,
}

var agentV1RunEvents = cli.Command{
	Name:    "events",
	Usage:   "Streams real-time run events over SSE. Supports replay using Last-Event-ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "last-event-id",
			Usage:     "Replay events after this sequence number",
			QueryPath: "lastEventId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentV1RunEvents,
	HideHelpCommand: true,
}

var agentV1RunExec = cli.Command{
	Name:    "exec",
	Usage:   "Starts execution of a queued run. The agent runs in a durable workflow — poll\n/run/:id/status for progress.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1RunExec,
	HideHelpCommand: true,
}

var agentV1RunGetDetails = cli.Command{
	Name:    "get-details",
	Usage:   "Full audit trail for a run including output, steps (tool calls, text), and token\nusage.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1RunGetDetails,
	HideHelpCommand: true,
}

var agentV1RunGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Lightweight status poll for a run. Use /run/:id/details for the full audit\ntrail.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1RunGetStatus,
	HideHelpCommand: true,
}

var agentV1RunWatch = cli.Command{
	Name:    "watch",
	Usage:   "Register a callback URL to receive notifications when the run completes. URL\nmust use https and must not point to a private network.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "HTTPS URL to receive completion callback",
			Required: true,
			BodyPath: "callbackUrl",
		},
	},
	Action:          handleAgentV1RunWatch,
	HideHelpCommand: true,
}

func handleAgentV1RunCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1RunNewParams{}

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
	_, err = client.Agent.V1.Run.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:run create", obj, format, transform)
}

func handleAgentV1RunList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1RunListParams{}

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
	_, err = client.Agent.V1.Run.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:run list", obj, format, transform)
}

func handleAgentV1RunCancel(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Run.Cancel(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:run cancel", obj, format, transform)
}

func handleAgentV1RunEvents(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1RunEventsParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.Agent.V1.Run.EventsStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(os.Stdout, "agent:v1:run events", stream, format, transform, maxItems)
}

func handleAgentV1RunExec(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Run.Exec(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:run exec", obj, format, transform)
}

func handleAgentV1RunGetDetails(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Run.GetDetails(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:run get-details", obj, format, transform)
}

func handleAgentV1RunGetStatus(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Run.GetStatus(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v1:run get-status", obj, format, transform)
}

func handleAgentV1RunWatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1RunWatchParams{}

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
	_, err = client.Agent.V1.Run.Watch(
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
	return ShowJSON(os.Stdout, "agent:v1:run watch", obj, format, transform)
}
