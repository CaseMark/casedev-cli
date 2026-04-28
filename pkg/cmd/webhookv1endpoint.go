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

var webhooksV1EndpointsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a webhook endpoint that receives platform events matching the supplied\nevent-type filters. Returns the generated signing secret ONCE — the response is\nthe only time it is shown in plaintext.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "event-type-filter",
			Usage:    `Glob patterns of event types to deliver (e.g. "vault.*", "ocr.job.completed", "*")`,
			Required: true,
			BodyPath: "eventTypeFilters",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "HTTPS callback URL that will receive event deliveries",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Human-readable label for this endpoint",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "resource-scopes",
			Usage:    "Optional per-resource allowlists. If vaultIds is set, only events for those vaults are delivered. Same for matterIds.",
			BodyPath: "resourceScopes",
		},
	},
	Action:          handleWebhooksV1EndpointsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"resource-scopes": {
		&requestflag.InnerFlag[[]string]{
			Name:       "resource-scopes.matter-ids",
			InnerField: "matterIds",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "resource-scopes.vault-ids",
			InnerField: "vaultIds",
		},
	},
})

var webhooksV1EndpointsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get webhook endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWebhooksV1EndpointsRetrieve,
	HideHelpCommand: true,
}

var webhooksV1EndpointsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially updates a webhook endpoint. Any omitted field is left unchanged.\nSigning secrets are rotated via the separate /rotate_secret endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type-filter",
			BodyPath: "eventTypeFilters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "resource-scopes",
			BodyPath: "resourceScopes",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "active", "disabled".`,
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			BodyPath: "url",
		},
	},
	Action:          handleWebhooksV1EndpointsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"resource-scopes": {
		&requestflag.InnerFlag[[]string]{
			Name:       "resource-scopes.matter-ids",
			InnerField: "matterIds",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "resource-scopes.vault-ids",
			InnerField: "vaultIds",
		},
	},
})

var webhooksV1EndpointsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the organization's webhook endpoints, newest first. Signing secrets are\nnever included.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by endpoint status",
			QueryPath: "status",
		},
	},
	Action:          handleWebhooksV1EndpointsList,
	HideHelpCommand: true,
}

var webhooksV1EndpointsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-deletes a webhook endpoint. Delivery stops immediately and the endpoint no\nlonger appears in list results. Delivery history is preserved (and can be\nfetched via GET /deliveries with the endpoint_id filter) so audit trails and\npost-mortem debugging remain possible.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWebhooksV1EndpointsDelete,
	HideHelpCommand: true,
}

var webhooksV1EndpointsRotateSecret = cli.Command{
	Name:    "rotate-secret",
	Usage:   "Generates a new signing secret for the endpoint. The previous secret remains\nvalid until `previousSecretExpiresInSec` elapses (default 24h, max 30 days).\nDuring the grace window deliveries are signed with both secrets so receivers can\nmigrate without downtime. Returns the new secret — this is the only time it is\nshown in plaintext.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "previous-secret-expires-in-sec",
			Usage:    "How long (seconds) the old secret continues to be accepted. 0 invalidates immediately. Default: 86400 (24h).",
			BodyPath: "previousSecretExpiresInSec",
		},
	},
	Action:          handleWebhooksV1EndpointsRotateSecret,
	HideHelpCommand: true,
}

var webhooksV1EndpointsTest = cli.Command{
	Name:    "test",
	Usage:   "Synchronously delivers a synthetic `webhook.test` event to the endpoint and\nreturns the HTTP result. No retries. Useful for validating that a new endpoint\nis reachable and its signature verifier works. The delivery is not persisted in\nthe delivery history.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "event-type",
			Usage:    `Event type to simulate. Defaults to "webhook.test".`,
			BodyPath: "eventType",
		},
		&requestflag.Flag[any]{
			Name:     "payload",
			Usage:    "Custom `data` payload. Defaults to a small placeholder.",
			BodyPath: "payload",
		},
	},
	Action:          handleWebhooksV1EndpointsTest,
	HideHelpCommand: true,
}

func handleWebhooksV1EndpointsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1EndpointNewParams{}

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

	return client.Webhooks.V1.Endpoints.New(ctx, params, options...)
}

func handleWebhooksV1EndpointsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Webhooks.V1.Endpoints.Get(ctx, cmd.Value("id").(string), options...)
}

func handleWebhooksV1EndpointsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1EndpointUpdateParams{}

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

	return client.Webhooks.V1.Endpoints.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleWebhooksV1EndpointsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1EndpointListParams{}

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

	return client.Webhooks.V1.Endpoints.List(ctx, params, options...)
}

func handleWebhooksV1EndpointsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Webhooks.V1.Endpoints.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleWebhooksV1EndpointsRotateSecret(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1EndpointRotateSecretParams{}

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

	return client.Webhooks.V1.Endpoints.RotateSecret(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleWebhooksV1EndpointsTest(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1EndpointTestParams{}

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

	return client.Webhooks.V1.Endpoints.Test(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
