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

var lincV1SessionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a Daytona-backed native Linc session with scoped Case.dev credentials.\nThis endpoint starts the sandbox actor only; messages and event replay use\nseparate endpoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "document-template-slug",
			Usage:    "Specific document template slugs to inject into the using-document-templates skill.",
			BodyPath: "documentTemplateSlugs",
		},
		&requestflag.Flag[*int64]{
			Name:     "idle-timeout-ms",
			BodyPath: "idleTimeoutMs",
		},
		&requestflag.Flag[*bool]{
			Name:     "include-document-templates",
			Usage:    "When true, inject all active org document templates into the using-document-templates skill.",
			BodyPath: "includeDocumentTemplates",
		},
		&requestflag.Flag[*string]{
			Name:     "instructions",
			Usage:    "Privileged C3-only hidden app instructions to append to the sandbox AGENTS.md.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[*string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[*string]{
			Name:     "scoped-api-key",
			Usage:    "Optional caller-provided scoped Case.dev API key for the runtime.",
			BodyPath: "scopedApiKey",
		},
		&requestflag.Flag[string]{
			Name:     "service-tier",
			Usage:    "Processing tier for eligible OpenAI GPT models. Priority provides lower latency at premium cost.",
			BodyPath: "serviceTier",
		},
		&requestflag.Flag[any]{
			Name:     "skill-slug",
			Usage:    "Skills API slugs to install into the runtime sandbox before the native session starts.",
			BodyPath: "skillSlugs",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			BodyPath: "title",
		},
		&requestflag.Flag[any]{
			Name:     "vault-id",
			BodyPath: "vaultIds",
		},
	},
	Action:          handleLincV1SessionsCreate,
	HideHelpCommand: true,
}

var lincV1SessionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "End native Linc session",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleLincV1SessionsDelete,
	HideHelpCommand: true,
}

var lincV1SessionsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Sends an abort RPC to the session runtime, ending the current turn while keeping\nthe session alive. Body handling is intentionally lenient — cancel is a stop\ncontrol, so unknown fields are ignored and an invalid or missing body is treated\nas empty rather than rejected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "clear-queue",
			Usage:    "Also clear queued steering/follow-up messages so the abort leaves the agent fully idle. Cleared texts are returned in the `response.data.clearedQueue` field of the response body. Without it, messages still queued when the abort settles are auto-continued as a new run. Runtimes older than the Linc release that supports this flag ignore it: the abort still happens but the queue is left untouched.",
			BodyPath: "clearQueue",
		},
	},
	Action:          handleLincV1SessionsCancel,
	HideHelpCommand: true,
}

var lincV1SessionsIngestEvents = requestflag.WithInnerFlags(cli.Command{
	Name:    "ingest-events",
	Usage:   "Runtime ingest endpoint for sandbox runtimes. Frames are persisted for replay;\nterminal frames emit the durable Linc session ended webhook.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "frame",
			Usage:    "Native Linc event frames to persist for replay.",
			Required: true,
			BodyPath: "frames",
		},
	},
	Action:          handleLincV1SessionsIngestEvents,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"frame": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "frame.event",
			Usage:      "Native Linc event payload.",
			InnerField: "event",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "frame.seq",
			Usage:      "Monotonic native event sequence number.",
			InnerField: "seq",
		},
		&requestflag.InnerFlag[string]{
			Name:       "frame.type",
			Usage:      "Native Linc event type.",
			InnerField: "type",
		},
	},
})

var lincV1SessionsRetrieveEvents = cli.Command{
	Name:    "retrieve-events",
	Usage:   "Returns persisted native Pi/Linc event envelopes after the requested cursor.\nLive delivery is handled by the Linc stream service.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "after-seq",
			Usage:     "Alias for cursor. Ignored when cursor is also provided.",
			QueryPath: "afterSeq",
		},
		&requestflag.Flag[int64]{
			Name:      "cursor",
			Usage:     "Replay events with a sequence number greater than this cursor.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[[]string]{
			Name:      "exclude-event-type",
			Usage:     "Comma-separated Linc event types to omit from replay.",
			QueryPath: "excludeEventTypes",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of events to return.",
			QueryPath: "limit",
		},
	},
	Action:          handleLincV1SessionsRetrieveEvents,
	HideHelpCommand: true,
}

var lincV1SessionsRetrieveMessages = cli.Command{
	Name:    "retrieve-messages",
	Usage:   "Returns completed Pi/Linc message entries derived from durable native Linc\nevents. This is the stable session-message read model for callers that need to\npersist or recover chat history without depending on a live SSE stream.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "after-seq",
			Usage:     "Alias for cursor. Ignored when cursor is also provided.",
			QueryPath: "afterSeq",
		},
		&requestflag.Flag[int64]{
			Name:      "cursor",
			Usage:     "Replay messages with a source event sequence number greater than this cursor.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of source events to scan for completed messages.",
			QueryPath: "limit",
		},
	},
	Action:          handleLincV1SessionsRetrieveMessages,
	HideHelpCommand: true,
}

var lincV1SessionsRetrieveState = cli.Command{
	Name:    "retrieve-state",
	Usage:   "Get native Linc session state",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleLincV1SessionsRetrieveState,
	HideHelpCommand: true,
}

var lincV1SessionsSendRpc = cli.Command{
	Name:    "send-rpc",
	Usage:   "Forwards a native Pi/Linc RPC command object to the sandbox-local Linc bridge\nunchanged. The route returns after Pi accepts or rejects the command; native\nevents are read through the events endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Native Pi/Linc RPC command type. Prompt commands also require a string id for idempotency.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Command idempotency key. Required when type is prompt.",
			BodyPath: "id",
		},
	},
	Action:          handleLincV1SessionsSendRpc,
	HideHelpCommand: true,
}

func handleLincV1SessionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := githubcomcasemarkcasedevgo.LincV1SessionNewParams{}

	return client.Linc.V1.Sessions.New(ctx, params, options...)
}

func handleLincV1SessionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Linc.V1.Sessions.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleLincV1SessionsCancel(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomcasemarkcasedevgo.LincV1SessionCancelParams{}

	return client.Linc.V1.Sessions.Cancel(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleLincV1SessionsIngestEvents(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomcasemarkcasedevgo.LincV1SessionIngestEventsParams{}

	return client.Linc.V1.Sessions.IngestEvents(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleLincV1SessionsRetrieveEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.LincV1SessionGetEventsParams{}

	return client.Linc.V1.Sessions.GetEvents(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleLincV1SessionsRetrieveMessages(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.LincV1SessionGetMessagesParams{}

	return client.Linc.V1.Sessions.GetMessages(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleLincV1SessionsRetrieveState(ctx context.Context, cmd *cli.Command) error {
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

	return client.Linc.V1.Sessions.GetState(ctx, cmd.Value("id").(string), options...)
}

func handleLincV1SessionsSendRpc(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomcasemarkcasedevgo.LincV1SessionSendRpcParams{}

	return client.Linc.V1.Sessions.SendRpc(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
