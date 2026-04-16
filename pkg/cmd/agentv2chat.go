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

var agentV2ChatCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a persistent OpenCode chat session backed by a Daytona runtime. Session\nstate is retained and can be resumed or recovered across requests.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "idle-timeout-ms",
			Usage:    "Idle timeout before the runtime is eligible to stop. Defaults to 15 minutes.",
			BodyPath: "idleTimeoutMs",
		},
		&requestflag.Flag[any]{
			Name:     "instructions",
			Usage:    "Optional hidden app instructions merged into the chat runtime bootstrap and never exposed as a user message. Only accepted for privileged C3 system keys.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Optional model override for the OpenCode session",
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
	Action:          handleAgentV2ChatCreate,
	HideHelpCommand: true,
}

var agentV2ChatDelete = cli.Command{
	Name:    "delete",
	Usage:   "Terminates the active Daytona runtime (if any), then marks the chat as ended.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV2ChatDelete,
	HideHelpCommand: true,
}

var agentV2ChatCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Aborts the active OpenCode generation for this Daytona-backed chat session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV2ChatCancel,
	HideHelpCommand: true,
}

var agentV2ChatCreateStreamToken = cli.Command{
	Name:    "create-stream-token",
	Usage:   "Returns a short-lived token that allows browser clients to connect directly to\nthe agent chat SSE stream without exposing the underlying org API key.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentV2ChatCreateStreamToken,
	HideHelpCommand: true,
}

var agentV2ChatReplyToQuestion = cli.Command{
	Name:    "reply-to-question",
	Usage:   "Answers a pending OpenCode question for the Daytona-backed chat session and\nresumes or recovers the runtime if needed.",
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
	Action:          handleAgentV2ChatReplyToQuestion,
	HideHelpCommand: true,
}

var agentV2ChatRespond = requestflag.WithInnerFlags(cli.Command{
	Name:    "respond",
	Usage:   "Streams a single assistant turn from a Daytona-backed chat runtime as normalized\nSSE events with stable turn, message, and part IDs. Emits events:\n`turn.started`, `turn.status`, `message.created`, `message.part.updated`,\n`message.completed`, `session.usage`, `turn.completed`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Optional model override. When provided, the runtime bootstrap config is updated so subsequent turns use this model. Conversation history is preserved.",
			BodyPath: "model",
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
	Action:          handleAgentV2ChatRespond,
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

var agentV2ChatSendMessage = requestflag.WithInnerFlags(cli.Command{
	Name:    "send-message",
	Usage:   "Sends a message to a Daytona-backed chat runtime and returns the complete\nresponse as a single JSON body. Blocks until the assistant turn completes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Optional model override. When provided, the runtime bootstrap config is updated so subsequent turns use this model. Conversation history is preserved.",
			BodyPath: "model",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "part",
			Usage:    `Message content parts. Currently only "text" type is supported. Additional types (e.g. file, image) may be added in future versions.`,
			BodyPath: "parts",
		},
	},
	Action:          handleAgentV2ChatSendMessage,
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

var agentV2ChatStream = cli.Command{
	Name:    "stream",
	Usage:   "Relays OpenCode SSE events for this Daytona-backed chat runtime. Supports replay\nfrom buffered events using Last-Event-ID and transparently reconnects stopped or\narchived runtimes. Accepts either Bearer token auth or a short-lived stream\ntoken via query parameter. When both are provided, Bearer auth takes precedence.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "token",
			Usage:     "Short-lived stream token from POST /agent/v2/chat/:id/stream-token. If provided, Bearer auth is not required.",
			QueryPath: "token",
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
	Action:          handleAgentV2ChatStream,
	HideHelpCommand: true,
}

func handleAgentV2ChatCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV2ChatNewParams{}

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
	_, err = client.Agent.V2.Chat.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v2:chat create", obj, format, transform)
}

func handleAgentV2ChatDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V2.Chat.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v2:chat delete", obj, format, transform)
}

func handleAgentV2ChatCancel(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V2.Chat.Cancel(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v2:chat cancel", obj, format, transform)
}

func handleAgentV2ChatCreateStreamToken(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.V2.Chat.NewStreamToken(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agent:v2:chat create-stream-token", obj, format, transform)
}

func handleAgentV2ChatReplyToQuestion(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.AgentV2ChatReplyToQuestionParams{}

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

	return client.Agent.V2.Chat.ReplyToQuestion(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("request-id").(string),
		params,
		options...,
	)
}

func handleAgentV2ChatRespond(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV2ChatRespondParams{}

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
	transform := cmd.Root().String("transform")
	stream := client.Agent.V2.Chat.RespondStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(os.Stdout, "agent:v2:chat respond", stream, format, transform, maxItems)
}

func handleAgentV2ChatSendMessage(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV2ChatSendMessageParams{}

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

	return client.Agent.V2.Chat.SendMessage(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleAgentV2ChatStream(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.AgentV2ChatStreamParams{}

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
	stream := client.Agent.V2.Chat.StreamStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(os.Stdout, "agent:v2:chat stream", stream, format, transform, maxItems)
}
