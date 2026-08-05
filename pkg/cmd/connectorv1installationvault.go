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

var connectorsV1InstallationsVaultsList = cli.Command{
	Name:    "list",
	Usage:   "List the vaults an installation may use, with capabilities and revocation state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleConnectorsV1InstallationsVaultsList,
	HideHelpCommand: true,
}

var connectorsV1InstallationsVaultsGrant = cli.Command{
	Name:    "grant",
	Usage:   "Grant (or update) an installation's access to a vault. Re-granting a revoked\nvault reactivates it. Import links need can_write; export links need can_read;\nmirror deletion and purge need can_manage.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "vault-id",
			Required:  true,
			PathParam: "vaultId",
		},
		&requestflag.Flag[bool]{
			Name:     "can-manage",
			Default:  false,
			BodyPath: "can_manage",
		},
		&requestflag.Flag[bool]{
			Name:     "can-read",
			Default:  true,
			BodyPath: "can_read",
		},
		&requestflag.Flag[bool]{
			Name:     "can-write",
			Default:  true,
			BodyPath: "can_write",
		},
		&requestflag.Flag[string]{
			Name:     "relationship",
			Usage:    `Allowed values: "owned", "shared".`,
			Default:  "owned",
			BodyPath: "relationship",
		},
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    `Allowed values: "provisioning", "lazy_reconcile", "explicit_share".`,
			Default:  "lazy_reconcile",
			BodyPath: "source",
		},
	},
	Action:          handleConnectorsV1InstallationsVaultsGrant,
	HideHelpCommand: true,
}

var connectorsV1InstallationsVaultsRevoke = cli.Command{
	Name:    "revoke",
	Usage:   "Revoke an installation's access to a vault. Links using the vault pause at their\nnext run; nothing is deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "vault-id",
			Required:  true,
			PathParam: "vaultId",
		},
	},
	Action:          handleConnectorsV1InstallationsVaultsRevoke,
	HideHelpCommand: true,
}

func handleConnectorsV1InstallationsVaultsList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Connectors.V1.Installations.Vaults.List(ctx, cmd.Value("id").(string), options...)
}

func handleConnectorsV1InstallationsVaultsGrant(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("vault-id") && len(unusedArgs) > 0 {
		cmd.Set("vault-id", unusedArgs[0])
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

	params := githubcomcasemarkcasedevgo.ConnectorV1InstallationVaultGrantParams{}

	return client.Connectors.V1.Installations.Vaults.Grant(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("vault-id").(string),
		params,
		options...,
	)
}

func handleConnectorsV1InstallationsVaultsRevoke(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("vault-id") && len(unusedArgs) > 0 {
		cmd.Set("vault-id", unusedArgs[0])
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

	return client.Connectors.V1.Installations.Vaults.Revoke(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("vault-id").(string),
		options...,
	)
}
