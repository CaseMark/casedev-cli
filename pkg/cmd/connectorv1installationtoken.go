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

var connectorsV1InstallationsTokensCreate = cli.Command{
	Name:    "create",
	Usage:   "Bound application management API key only. Returns a 256-bit opaque bearer\nsecret once; valid for 15 minutes, connector data-plane only. No refresh token.\nAt most 4 overlapping credentials and 20 mints/minute per installation. The\napplication server must derive tenant identity; this is not browser\nauthentication.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "scope",
			Required: true,
			BodyPath: "scopes",
		},
	},
	Action:          handleConnectorsV1InstallationsTokensCreate,
	HideHelpCommand: true,
}

var connectorsV1InstallationsTokensRevoke = cli.Command{
	Name:    "revoke",
	Usage:   "Issuing API key only; installation credentials cannot revoke or mint\ncredentials. Immediate revocation on the next request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "token-id",
			Required:  true,
			PathParam: "tokenId",
		},
	},
	Action:          handleConnectorsV1InstallationsTokensRevoke,
	HideHelpCommand: true,
}

func handleConnectorsV1InstallationsTokensCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1InstallationTokenNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Installations.Tokens.New(
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
		Title:          "connectors:v1:installations:tokens create",
		Transform:      transform,
	})
}

func handleConnectorsV1InstallationsTokensRevoke(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("token-id") && len(unusedArgs) > 0 {
		cmd.Set("token-id", unusedArgs[0])
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

	return client.Connectors.V1.Installations.Tokens.Revoke(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("token-id").(string),
		options...,
	)
}
