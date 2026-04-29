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

var webhooksV1DeliveriesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get webhook delivery",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWebhooksV1DeliveriesRetrieve,
	HideHelpCommand: true,
}

var webhooksV1DeliveriesList = cli.Command{
	Name:    "list",
	Usage:   "Returns delivery attempts for the organization, newest first. Filter by\nendpoint_id or status to narrow results.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "endpoint-id",
			QueryPath: "endpoint_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pending", "delivered", "failed".`,
			QueryPath: "status",
		},
	},
	Action:          handleWebhooksV1DeliveriesList,
	HideHelpCommand: true,
}

var webhooksV1DeliveriesReplay = cli.Command{
	Name:    "replay",
	Usage:   "Re-sends the original event to its endpoint. The payload is reconstructed from\nthe delivery record (same eventId, eventType, and occurred_at). Replay\ndeliveries include a Case.dev replay marker header so receivers can distinguish\nreplays from first-time deliveries. Uses the endpoint's current signing secret —\nnot the one in force at the original delivery time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "payload",
			Usage:    "Override payload to deliver. Must only be supplied when the delivery record lacks enough context to reconstruct the original event (rare). Defaults to an empty data envelope.",
			BodyPath: "payload",
		},
	},
	Action:          handleWebhooksV1DeliveriesReplay,
	HideHelpCommand: true,
}

func handleWebhooksV1DeliveriesRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Webhooks.V1.Deliveries.Get(ctx, cmd.Value("id").(string), options...)
}

func handleWebhooksV1DeliveriesList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1DeliveryListParams{}

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

	return client.Webhooks.V1.Deliveries.List(ctx, params, options...)
}

func handleWebhooksV1DeliveriesReplay(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.WebhookV1DeliveryReplayParams{}

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

	return client.Webhooks.V1.Deliveries.Replay(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
