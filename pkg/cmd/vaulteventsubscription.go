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

var vaultEventsSubscriptionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a webhook subscription for vault lifecycle events. Optional object\nfilters can limit notifications to specific vault objects.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "Webhook endpoint URL that will receive vault event deliveries",
			Required: true,
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			Usage:    "Vault event types to deliver. Omit to receive the default supported set.",
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[[]string]{
			Name:     "object-id",
			Usage:    "Vault object IDs to limit notifications to. Omit to receive events for all objects in the vault.",
			BodyPath: "objectIds",
		},
		&requestflag.Flag[string]{
			Name:     "signing-secret",
			Usage:    "Optional secret used to sign outbound webhook deliveries",
			BodyPath: "signingSecret",
		},
	},
	Action:          handleVaultEventsSubscriptionsCreate,
	HideHelpCommand: true,
}

var vaultEventsSubscriptionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates callback URL, filters, active state, or signing secret for a vault\nwebhook subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "Updated webhook endpoint URL for deliveries",
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[bool]{
			Name:     "clear-signing-secret",
			Usage:    "Whether to remove the existing signing secret",
			BodyPath: "clearSigningSecret",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			Usage:    "Updated event types to deliver for this subscription",
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			Usage:    "Whether the subscription should continue delivering events",
			BodyPath: "isActive",
		},
		&requestflag.Flag[[]string]{
			Name:     "object-id",
			Usage:    "Updated vault object IDs to limit notifications to. Pass an empty array to remove the filter.",
			BodyPath: "objectIds",
		},
		&requestflag.Flag[string]{
			Name:     "signing-secret",
			Usage:    "Replacement secret used to sign webhook deliveries",
			BodyPath: "signingSecret",
		},
	},
	Action:          handleVaultEventsSubscriptionsUpdate,
	HideHelpCommand: true,
}

var vaultEventsSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Lists webhook subscriptions configured for a vault.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVaultEventsSubscriptionsList,
	HideHelpCommand: true,
}

var vaultEventsSubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deactivates a vault webhook subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
	},
	Action:          handleVaultEventsSubscriptionsDelete,
	HideHelpCommand: true,
}

var vaultEventsSubscriptionsTest = cli.Command{
	Name:    "test",
	Usage:   "Delivers a test event to a single vault webhook subscription. Uses the same\npayload shape, signature, and retry behavior as production event delivery.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "event-type",
			Usage:    "Optional event type override for this test",
			BodyPath: "eventType",
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Usage:    "Optional object ID for object-scoped payload testing",
			BodyPath: "objectId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "payload",
			Usage:    "Optional additional fields merged into payload.data",
			BodyPath: "payload",
		},
	},
	Action:          handleVaultEventsSubscriptionsTest,
	HideHelpCommand: true,
}

func handleVaultEventsSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultEventSubscriptionNewParams{}

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

	return client.Vault.Events.Subscriptions.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleVaultEventsSubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultEventSubscriptionUpdateParams{}

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

	return client.Vault.Events.Subscriptions.Update(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("subscription-id").(string),
		params,
		options...,
	)
}

func handleVaultEventsSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Vault.Events.Subscriptions.List(ctx, cmd.Value("id").(string), options...)
}

func handleVaultEventsSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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

	return client.Vault.Events.Subscriptions.Delete(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("subscription-id").(string),
		options...,
	)
}

func handleVaultEventsSubscriptionsTest(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultEventSubscriptionTestParams{}

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

	return client.Vault.Events.Subscriptions.Test(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("subscription-id").(string),
		params,
		options...,
	)
}
