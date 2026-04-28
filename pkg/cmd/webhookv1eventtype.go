// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-go"
	"github.com/urfave/cli/v3"
)

var webhooksV1EventTypesList = cli.Command{
	Name:            "list",
	Usage:           "Returns the catalog of event types that can be subscribed to via webhook\nendpoints. Each entry lists the required service scope the API key must carry to\nsubscribe, plus the stability level.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWebhooksV1EventTypesList,
	HideHelpCommand: true,
}

func handleWebhooksV1EventTypesList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Webhooks.V1.EventTypes.List(ctx, options...)
}
