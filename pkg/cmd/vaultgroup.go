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

var vaultGroupsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a vault group for organizing vaults and applying group-scoped access\ncontrols. Group-scoped API keys cannot create or manage vault groups.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the vault group",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description of the vault group purpose",
			BodyPath: "description",
		},
	},
	Action:          handleVaultGroupsCreate,
	HideHelpCommand: true,
}

var vaultGroupsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a vault group for the authenticated organization. Only provided fields\nare changed, and setting description to null removes the current description.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Updated vault group description. Pass null to remove the current description.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New human-readable name for the vault group",
			BodyPath: "name",
		},
	},
	Action:          handleVaultGroupsUpdate,
	HideHelpCommand: true,
}

var vaultGroupsList = cli.Command{
	Name:            "list",
	Usage:           "Lists vault groups visible to the authenticated organization. Group-scoped API\nkeys only receive groups within their allowed scope.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleVaultGroupsList,
	HideHelpCommand: true,
}

var vaultGroupsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-deletes a vault group that no longer has any active vaults assigned. This\noperation is blocked when the group still contains vaults.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group-id",
			Required: true,
		},
	},
	Action:          handleVaultGroupsDelete,
	HideHelpCommand: true,
}

func handleVaultGroupsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultGroupNewParams{}

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

	return client.Vault.Groups.New(ctx, params, options...)
}

func handleVaultGroupsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group-id") && len(unusedArgs) > 0 {
		cmd.Set("group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultGroupUpdateParams{}

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

	return client.Vault.Groups.Update(
		ctx,
		cmd.Value("group-id").(string),
		params,
		options...,
	)
}

func handleVaultGroupsList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Vault.Groups.List(ctx, options...)
}

func handleVaultGroupsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group-id") && len(unusedArgs) > 0 {
		cmd.Set("group-id", unusedArgs[0])
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

	return client.Vault.Groups.Delete(ctx, cmd.Value("group-id").(string), options...)
}
