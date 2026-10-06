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

var mattersV1ContentPurgesCreate = cli.Command{
	Name:    "create",
	Usage:   "Queues an idempotent hard deletion of explicitly owned content while preserving\nthe Matter, Vault, and unrelated content. Use unified matter.content_purge\nwebhooks for completion, not polling.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "request-id",
			Usage:    "Stable caller idempotency ID; cannot be reused with different targets.",
			Required: true,
			BodyPath: "request_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "object-id",
			BodyPath: "object_ids",
		},
		&requestflag.Flag[[]string]{
			Name:     "session-id",
			BodyPath: "session_ids",
		},
		&requestflag.Flag[[]string]{
			Name:     "transcription-id",
			BodyPath: "transcription_ids",
		},
		&requestflag.Flag[[]string]{
			Name:     "work-item-id",
			BodyPath: "work_item_ids",
		},
	},
	Action:          handleMattersV1ContentPurgesCreate,
	HideHelpCommand: true,
}

var mattersV1ContentPurgesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Owner-only receipt for operator diagnostics. Integrations must use unified\ncontent-purge webhooks rather than polling.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "purge-id",
			Required:  true,
			PathParam: "purgeId",
		},
	},
	Action:          handleMattersV1ContentPurgesRetrieve,
	HideHelpCommand: true,
}

func handleMattersV1ContentPurgesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.MatterV1ContentPurgeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Matters.V1.ContentPurges.New(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "matters:v1:content-purges create",
		Transform:      transform,
	})
}

func handleMattersV1ContentPurgesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("purge-id") && len(unusedArgs) > 0 {
		cmd.Set("purge-id", unusedArgs[0])
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
	_, err = client.Matters.V1.ContentPurges.Get(ctx, cmd.Value("purge-id").(string), options...)
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
		Title:          "matters:v1:content-purges retrieve",
		Transform:      transform,
	})
}
