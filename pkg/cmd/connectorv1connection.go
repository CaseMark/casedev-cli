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

var connectorsV1ConnectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a pending provider connection and return a one-time connect_url for the\nhosted OAuth flow. The user completes provider consent at connect_url and is\nredirected to return_url with ?connection_id=.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    `Allowed values: "clio", "gdrive", "microsoft".`,
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			Usage:    "HTTPS URL the user is sent back to after consent.",
			Required: true,
			BodyPath: "return_url",
		},
		&requestflag.Flag[string]{
			Name:     "scope-tier",
			Usage:    "Provider-specific OAuth permission tier. Omit to use the provider's default.",
			BodyPath: "scope_tier",
		},
	},
	Action:          handleConnectorsV1ConnectionsCreate,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve one provider connection, including account identity and health.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleConnectorsV1ConnectionsRetrieve,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsList = cli.Command{
	Name:    "list",
	Usage:   "List provider connections for the organization, with health status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "provider",
			QueryPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pending", "healthy", "reauth_required", "revoked", "throttled".`,
			QueryPath: "status",
		},
	},
	Action:          handleConnectorsV1ConnectionsList,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Unlink a provider account: revoke tokens at the provider and delete them.\npurge=true additionally deletes the vault documents its import links brought in.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:      "purge",
			Default:   false,
			QueryPath: "purge",
		},
	},
	Action:          handleConnectorsV1ConnectionsDelete,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsBrowse = cli.Command{
	Name:    "browse",
	Usage:   "Browse the provider one level at a time. Without a site, container, or parent,\nreturns top-level resources. Pass the stable browse_ref fields returned by one\nresponse to navigate into the next level. Returns 403\nprovider_scope_insufficient when the connection scope cannot browse server-side.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "container",
			Usage:     "Container id to list, or the container containing parent",
			QueryPath: "container",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "parent",
			Usage:     "Folder id to list",
			QueryPath: "parent",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Optional provider-supported search text",
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "site",
			Usage:     "Site id to list",
			QueryPath: "site",
		},
	},
	Action:          handleConnectorsV1ConnectionsBrowse,
	HideHelpCommand: true,
}

func handleConnectorsV1ConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.New(ctx, params, options...)
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
		Title:          "connectors:v1:connections create",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Connectors.V1.Connections.Get(ctx, cmd.Value("id").(string), options...)
}

func handleConnectorsV1ConnectionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.List(ctx, params, options...)
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
		Title:          "connectors:v1:connections list",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionDeleteParams{}

	return client.Connectors.V1.Connections.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleConnectorsV1ConnectionsBrowse(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionBrowseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.Browse(
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
		Title:          "connectors:v1:connections browse",
		Transform:      transform,
	})
}
