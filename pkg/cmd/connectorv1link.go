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

var connectorsV1LinksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve one link: state, counts, and embedded active_run/last_run. Poll this\nafter POST /transfer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleConnectorsV1LinksRetrieve,
	HideHelpCommand: true,
}

var connectorsV1LinksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Pause/resume a link (state \"paused\" | \"ready\"), change its mode (synced -> once\nis the sync downgrade), or edit its policy in place.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    `Allowed values: "once", "synced".`,
			BodyPath: "mode",
		},
		&requestflag.Flag[any]{
			Name:     "policy",
			BodyPath: "policy",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    `Allowed values: "paused", "ready".`,
			BodyPath: "state",
		},
	},
	Action:          handleConnectorsV1LinksUpdate,
	HideHelpCommand: true,
}

var connectorsV1LinksList = cli.Command{
	Name:    "list",
	Usage:   "List transfer links, filterable by vault, connection, direction, mode, and\nstate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "connection-id",
			QueryPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:      "direction",
			Usage:     `Allowed values: "import", "export".`,
			QueryPath: "direction",
		},
		&requestflag.Flag[string]{
			Name:      "mode",
			Usage:     `Allowed values: "once", "synced".`,
			QueryPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:      "pair-id",
			QueryPath: "pair_id",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     `Allowed values: "ready", "running", "active", "paused", "orphaned", "error".`,
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "vault-id",
			QueryPath: "vault_id",
		},
	},
	Action:          handleConnectorsV1LinksList,
	HideHelpCommand: true,
}

var connectorsV1LinksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a link and its ledger. vault_docs=delete additionally removes the vault\ndocuments an import link brought in (default: keep).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "vault-docs",
			Usage:     `Allowed values: "keep", "delete".`,
			Default:   "keep",
			QueryPath: "vault_docs",
		},
	},
	Action:          handleConnectorsV1LinksDelete,
	HideHelpCommand: true,
}

var connectorsV1LinksListObjects = cli.Command{
	Name:    "list-objects",
	Usage:   "Per-file transfer ledger for a link: provider item, vault object, path, content\nversion, state, and error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     `Allowed values: "pending", "transferring", "ingesting", "synced", "skipped", "failed", "tombstoned".`,
			QueryPath: "state",
		},
	},
	Action:          handleConnectorsV1LinksListObjects,
	HideHelpCommand: true,
}

func handleConnectorsV1LinksRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Connectors.V1.Links.Get(ctx, cmd.Value("id").(string), options...)
}

func handleConnectorsV1LinksUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1LinkUpdateParams{}

	return client.Connectors.V1.Links.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleConnectorsV1LinksList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1LinkListParams{}

	return client.Connectors.V1.Links.List(ctx, params, options...)
}

func handleConnectorsV1LinksDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1LinkDeleteParams{}

	return client.Connectors.V1.Links.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleConnectorsV1LinksListObjects(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1LinkListObjectsParams{}

	return client.Connectors.V1.Links.ListObjects(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
