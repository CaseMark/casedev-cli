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

var connectorsV1Fetch = cli.Command{
	Name:    "fetch",
	Usage:   "Copy one provider file, named by an open_ref from browse, into a Vault\n(vault_id) or to a presigned S3 PUT URL (upload_url). Runs synchronously; files\nover 100 MB return 413 and belong in a transfer. A repeat for an unchanged\nprovider version returns the first copy with 200. Vault copies are stored as\nuploaded: call POST /vault/{id}/ingest/{objectId} to index them. Fetched copies\nare never exported by a link.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "open-ref",
			Usage:    "From a browse item, exactly as received",
			Required: true,
			BodyPath: "open_ref",
		},
		&requestflag.Flag[string]{
			Name:     "upload-url",
			Usage:    "Presigned HTTPS S3 PUT URL. Content-Type must not be a signed header. Exactly one destination.",
			BodyPath: "upload_url",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Usage:    "Destination Vault. Exactly one destination.",
			BodyPath: "vault_id",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1Fetch,
	HideHelpCommand: true,
}

var connectorsV1SyncLink = requestflag.WithInnerFlags(cli.Command{
	Name:    "sync-link",
	Usage:   "Standing promise: backfill now, then stay current (the sync sweeper re-runs\nsynced links on a schedule). Direction both creates paired import/export links\nand defaults export to a CaseMark Output subfolder. Same body as /transfer minus\nrun_mode. Upserts links by (connection_id, direction, remote, vault_id,\nselection); existing once-links are upgraded in place with their ledger and\ncursor preserved. Downgrade or pause via PATCH /links/{id}.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "direction",
			Usage:    `Allowed values: "import", "export", "both".`,
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
		&requestflag.Flag[map[string]any]{
			Name:     "export-destination",
			Usage:    "Optional destination for direction both. Defaults to CaseMark Output under remote.",
			BodyPath: "export_destination",
		},
		&requestflag.Flag[*string]{
			Name:     "matter-id",
			BodyPath: "matter_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "policy",
			BodyPath: "policy",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "selection",
			Usage:    "Export only part of the vault (direction export only; 400 selection_requires_export otherwise). Provide object_ids, folder, or both (their intersection). The selection is part of link identity: each distinct selection to the same destination is its own link, and repeating a selection re-runs that link. Exported objects still exclude files still processing and connector-imported or fetched copies. Destination paths: with folder, paths are relative to that folder, so its contents land directly in remote; with object_ids only, each object keeps its full vault path. Omit to export the whole vault.",
			BodyPath: "selection",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
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
			Name:       "remote.resource-type",
			InnerField: "resource_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.site-id",
			InnerField: "site_id",
		},
	},
	"export-destination": {
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.folder-id",
			InnerField: "folder_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.container-id",
			InnerField: "container_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.site-id",
			InnerField: "site_id",
		},
	},
	"policy": {
		&requestflag.InnerFlag[string]{
			Name:       "policy.collisions",
			Usage:      "Export handling of a destination file with the same name that this export does not own. version keeps both: the export is written as name (CaseMark).ext, then name (CaseMark 2).ext, name (CaseMark 3).ext, using the first number the folder does not use (up to 100). overwrite replaces the existing file; skip leaves it. Any other value returns 400.",
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
	"selection": {
		&requestflag.InnerFlag[string]{
			Name:       "selection.folder",
			Usage:      "Vault folder path such as /Reports/Final; includes every subfolder. The vault root is not a selection.",
			InnerField: "folder",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "selection.object-ids",
			Usage:      "Vault object ids from vault_id. Unknown ids return 400 selection_objects_not_found with data.error.count.",
			InnerField: "object_ids",
		},
	},
})

var connectorsV1Transfer = requestflag.WithInnerFlags(cli.Command{
	Name:    "transfer",
	Usage:   "One-shot import (provider folder → vault), export (vault → provider folder), or\nboth. Direction both creates paired import/export links and defaults export to a\nCaseMark Output subfolder. Upserts links by (connection_id, direction, remote,\nvault_id, selection): first call backfills, later calls move only new/changed\nfiles via the ledger. Poll GET /links/{id} → active_run for progress.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "direction",
			Usage:    `Allowed values: "import", "export", "both".`,
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
		&requestflag.Flag[map[string]any]{
			Name:     "export-destination",
			Usage:    "Optional destination for direction both. Defaults to CaseMark Output under remote.",
			BodyPath: "export_destination",
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
		&requestflag.Flag[map[string]any]{
			Name:     "selection",
			Usage:    "Export only part of the vault (direction export only; 400 selection_requires_export otherwise). Provide object_ids, folder, or both (their intersection). The selection is part of link identity: each distinct selection to the same destination is its own link, and repeating a selection re-runs that link. Exported objects still exclude files still processing and connector-imported or fetched copies. Destination paths: with folder, paths are relative to that folder, so its contents land directly in remote; with object_ids only, each object keeps its full vault path. Omit to export the whole vault.",
			BodyPath: "selection",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
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
			Name:       "remote.resource-type",
			InnerField: "resource_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "remote.site-id",
			InnerField: "site_id",
		},
	},
	"export-destination": {
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.folder-id",
			InnerField: "folder_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.container-id",
			InnerField: "container_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "export-destination.site-id",
			InnerField: "site_id",
		},
	},
	"policy": {
		&requestflag.InnerFlag[string]{
			Name:       "policy.collisions",
			Usage:      "Export handling of a destination file with the same name that this export does not own. version keeps both: the export is written as name (CaseMark).ext, then name (CaseMark 2).ext, name (CaseMark 3).ext, using the first number the folder does not use (up to 100). overwrite replaces the existing file; skip leaves it. Any other value returns 400.",
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
	"selection": {
		&requestflag.InnerFlag[string]{
			Name:       "selection.folder",
			Usage:      "Vault folder path such as /Reports/Final; includes every subfolder. The vault root is not a selection.",
			InnerField: "folder",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "selection.object-ids",
			Usage:      "Vault object ids from vault_id. Unknown ids return 400 selection_objects_not_found with data.error.count.",
			InnerField: "object_ids",
		},
	},
})

func handleConnectorsV1Fetch(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1FetchParams{}

	return client.Connectors.V1.Fetch(ctx, params, options...)
}

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
