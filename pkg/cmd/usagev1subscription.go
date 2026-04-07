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

var usageV1SubscriptionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a webhook subscription for usage, balance, and billing events.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Required: true,
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[string]{
			Name:     "signing-secret",
			BodyPath: "signingSecret",
		},
	},
	Action:          handleUsageV1SubscriptionsCreate,
	HideHelpCommand: true,
}

var usageV1SubscriptionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates callback URL, event filters, active state, or signing secret.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[bool]{
			Name:     "clear-signing-secret",
			BodyPath: "clearSigningSecret",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "isActive",
		},
		&requestflag.Flag[string]{
			Name:     "signing-secret",
			BodyPath: "signingSecret",
		},
	},
	Action:          handleUsageV1SubscriptionsUpdate,
	HideHelpCommand: true,
}

var usageV1SubscriptionsList = cli.Command{
	Name:            "list",
	Usage:           "Lists webhook subscriptions configured for usage and billing events.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUsageV1SubscriptionsList,
	HideHelpCommand: true,
}

var usageV1SubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deactivates a usage webhook subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
	},
	Action:          handleUsageV1SubscriptionsDelete,
	HideHelpCommand: true,
}

var usageV1SubscriptionsTest = cli.Command{
	Name:    "test",
	Usage:   "Delivers a test event to a single usage webhook subscription using the same\npayload shape and signing behavior as production delivery.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "event-type",
			BodyPath: "eventType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "payload",
			BodyPath: "payload",
		},
	},
	Action:          handleUsageV1SubscriptionsTest,
	HideHelpCommand: true,
}

func handleUsageV1SubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.UsageV1SubscriptionNewParams{}

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

	return client.Usage.V1.Subscriptions.New(ctx, params, options...)
}

func handleUsageV1SubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.UsageV1SubscriptionUpdateParams{}

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

	return client.Usage.V1.Subscriptions.Update(
		ctx,
		cmd.Value("subscription-id").(string),
		params,
		options...,
	)
}

func handleUsageV1SubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	return client.Usage.V1.Subscriptions.List(ctx, options...)
}

func handleUsageV1SubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
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

	return client.Usage.V1.Subscriptions.Delete(ctx, cmd.Value("subscription-id").(string), options...)
}

func handleUsageV1SubscriptionsTest(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.UsageV1SubscriptionTestParams{}

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

	return client.Usage.V1.Subscriptions.Test(
		ctx,
		cmd.Value("subscription-id").(string),
		params,
		options...,
	)
}
