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

var connectorsV1SyncLink = requestflag.WithInnerFlags(cli.Command{
	Name:    "sync-link",
	Usage:   "Standing promise: backfill now, then stay current (the sync sweeper re-runs\nsynced links on a schedule). Same body as /transfer minus run_mode. Upserts the\nlink identified by (connection_id, direction, remote, vault_id); an existing\nonce-link is upgraded in place with its ledger and cursor preserved. Downgrade\nor pause via PATCH /links/{id}.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "direction",
			Usage:    `Allowed values: "import", "export".`,
			Required: true,
			BodyPath: "direction",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "remote",
			Required: true,
			BodyPath: "remote",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Required: true,
			BodyPath: "vault_id",
		},
		&requestflag.Flag[*string]{
			Name:     "matter-id",
			BodyPath: "matter_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "policy",
			BodyPath: "policy",
		},
	},
	Action:          handleConnectorsV1SyncLink,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"remote": {
		&requestflag.InnerFlag[string]{
			Name:       "remote.folder-id",
			InnerField: "folder_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.container-id",
			InnerField: "container_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.site-id",
			InnerField: "site_id",
		},
	},
	"policy": {
		&requestflag.InnerFlag[string]{
			Name:       "policy.collisions",
			Usage:      `Allowed values: "version", "overwrite", "skip".`,
			InnerField: "collisions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "policy.deletes",
			Usage:      `Allowed values: "mirror", "preserve".`,
			InnerField: "deletes",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "policy.filters",
			InnerField: "filters",
		},
	},
})

var connectorsV1Transfer = requestflag.WithInnerFlags(cli.Command{
	Name:    "transfer",
	Usage:   "One-shot import (provider folder → vault) or export (vault → provider folder).\nUpserts the link identified by (connection_id, direction, remote, vault_id):\nfirst call backfills, later calls move only new/changed files via the ledger.\nPoll GET /links/{id} → active_run for progress.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "direction",
			Usage:    `Allowed values: "import", "export".`,
			Required: true,
			BodyPath: "direction",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "remote",
			Required: true,
			BodyPath: "remote",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Required: true,
			BodyPath: "vault_id",
		},
		&requestflag.Flag[*string]{
			Name:     "matter-id",
			BodyPath: "matter_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "policy",
			BodyPath: "policy",
		},
		&requestflag.Flag[string]{
			Name:     "run-mode",
			Usage:    `Allowed values: "auto", "full_reconcile".`,
			Default:  "auto",
			BodyPath: "run_mode",
		},
	},
	Action:          handleConnectorsV1Transfer,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"remote": {
		&requestflag.InnerFlag[string]{
			Name:       "remote.folder-id",
			InnerField: "folder_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.container-id",
			InnerField: "container_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.site-id",
			InnerField: "site_id",
		},
	},
	"policy": {
		&requestflag.InnerFlag[string]{
			Name:       "policy.collisions",
			Usage:      `Allowed values: "version", "overwrite", "skip".`,
			InnerField: "collisions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "policy.deletes",
			Usage:      `Allowed values: "mirror", "preserve".`,
			InnerField: "deletes",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "policy.filters",
			InnerField: "filters",
		},
	},
})

func handleConnectorsV1SyncLink(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.SyncLink(ctx, params, options...)
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
		Title:          "connectors:v1 sync-link",
		Transform:      transform,
	})
}

func handleConnectorsV1Transfer(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1TransferParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Transfer(ctx, params, options...)
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
		Title:          "connectors:v1 transfer",
		Transform:      transform,
	})
}
