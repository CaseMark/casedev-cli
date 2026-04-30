// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var agentV1ChatCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a persistent chat session backed by a Daytona or Vercel runtime. Session\nstate is retained and can be resumed or recovered across requests.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "idle-timeout-ms",
			Usage:    "Idle timeout before session is eligible for snapshot/termination. Defaults to 15 minutes.",
			BodyPath: "idleTimeoutMs",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Optional model override for the chat runtime session",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    "Optional human-readable session title",
			BodyPath: "title",
		},
		&requestflag.Flag[any]{
			Name:     "vault-id",
			Usage:    "Restrict the chat session to specific vault IDs",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleAgentV1ChatCreate,
	HideHelpCommand: true,
}

var agentV1ChatDelete = cli.Command{
	Name:    "delete",
	Usage:   "Snapshots and terminates the active sandbox (if any), then marks the chat as\nended.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1ChatDelete,
	HideHelpCommand: true,
}

var agentV1ChatCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Aborts the active generation for this chat session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV1ChatCancel,
	HideHelpCommand: true,
}

var agentV1ChatReplyToQuestion = cli.Command{
	Name:    "reply-to-question",
	Usage:   "Answers a pending runtime question for the chat session bound to this agent\nchat.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "request-id",
			Required: true,
		},
		&requestflag.Flag[[]any]{
			Name:     "answer",
			Usage:    "Answer selections for each prompt element in the pending question",
			Required: true,
			BodyPath: "answers",
		},
	},
	Action:          handleAgentV1ChatReplyToQuestion,
	HideHelpCommand: true,
}

var agentV1ChatRespond = requestflag.WithInnerFlags(cli.Command{
	Name:    "respond",
	Usage:   "Streams a single assistant turn as normalized SSE events with stable turn,\nmessage, and part IDs. Emits events: `turn.started`, `turn.status`,\n`message.created`, `message.part.updated`, `message.completed`, `session.usage`,\n`turn.completed`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "part",
			Usage:    `Message content parts. Currently only "text" type is supported. Additional types (e.g. file, image) may be added in future versions.`,
			BodyPath: "parts",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentV1ChatRespond,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"part": {
		&requestflag.InnerFlag[string]{
			Name:       "part.text",
			Usage:      "The message text content",
			InnerField: "text",
		},
		&requestflag.InnerFlag[string]{
			Name:       "part.type",
			Usage:      `Part type. Currently only "text" is supported.`,
			InnerField: "type",
		},
	},
})

var agentV1ChatSendMessage = requestflag.WithInnerFlags(cli.Command{
	Name:    "send-message",
	Usage:   "Sends a message and returns the complete response as a single JSON body. Blocks\nuntil the agent turn completes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "part",
			Usage:    `Message content parts. Currently only "text" type is supported. Additional types (e.g. file, image) may be added in future versions.`,
			BodyPath: "parts",
		},
	},
	Action:          handleAgentV1ChatSendMessage,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"part": {
		&requestflag.InnerFlag[string]{
			Name:       "part.text",
			Usage:      "The message text content",
			InnerField: "text",
		},
		&requestflag.InnerFlag[string]{
			Name:       "part.type",
			Usage:      `Part type. Currently only "text" is supported.`,
			InnerField: "type",
		},
	},
})

var agentV1ChatStream = cli.Command{
	Name:    "stream",
	Usage:   "Relays runtime SSE events for this chat. Supports replay from buffered events\nusing Last-Event-ID.",
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
	Action:          handleAgentV1ChatStream,
	HideHelpCommand: true,
}

func handleAgentV1ChatCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1ChatNewParams{}

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
	_, err = client.Agent.V1.Chat.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent:v1:chat create",
		Transform:      transform,
	})
}

func handleAgentV1ChatDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Chat.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent:v1:chat delete",
		Transform:      transform,
	})
}

func handleAgentV1ChatCancel(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V1.Chat.Cancel(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent:v1:chat cancel",
		Transform:      transform,
	})
}

func handleAgentV1ChatReplyToQuestion(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("request-id") && len(unusedArgs) > 0 {
		cmd.Set("request-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1ChatReplyToQuestionParams{}

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

	return client.Agent.V1.Chat.ReplyToQuestion(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("request-id").(string),
		params,
		options...,
	)
}

func handleAgentV1ChatRespond(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1ChatRespondParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	stream := client.Agent.V1.Chat.RespondStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(stream, maxItems, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent:v1:chat respond",
		Transform:      transform,
	})
}

func handleAgentV1ChatSendMessage(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1ChatSendMessageParams{}

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

	return client.Agent.V1.Chat.SendMessage(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleAgentV1ChatStream(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV1ChatStreamParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	stream := client.Agent.V1.Chat.StreamStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(stream, maxItems, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent:v1:chat stream",
		Transform:      transform,
	})
}
