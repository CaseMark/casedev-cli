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

var vaultMemoryCreate = cli.Command{
	Name:    "create",
	Usage:   "Append a new file-backed memory entry to a vault.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "fact", "party", "issue", "deadline", "discovery", "correction", "preference".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "source",
			BodyPath: "source",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleVaultMemoryCreate,
	HideHelpCommand: true,
}

var vaultMemoryUpdate = cli.Command{
	Name:    "update",
	Usage:   "Rewrite a file-backed vault memory entry with updated content, source, or tags.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "entry-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content",
			BodyPath: "content",
		},
		&requestflag.Flag[any]{
			Name:     "source",
			BodyPath: "source",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleVaultMemoryUpdate,
	HideHelpCommand: true,
}

var vaultMemoryList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve file-backed memory entries stored in a vault.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVaultMemoryList,
	HideHelpCommand: true,
}

var vaultMemoryDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a file-backed memory entry from a vault.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "entry-id",
			Required: true,
		},
	},
	Action:          handleVaultMemoryDelete,
	HideHelpCommand: true,
}

var vaultMemorySearch = cli.Command{
	Name:    "search",
	Usage:   "Search file-backed vault memory using simple full-text matching over content and\ntags.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Default:  10,
			BodyPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[[]string]{
			Name:     "type",
			BodyPath: "types",
		},
	},
	Action:          handleVaultMemorySearch,
	HideHelpCommand: true,
}

func handleVaultMemoryCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultMemoryNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.Memory.New(
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
		Title:          "vault:memory create",
		Transform:      transform,
	})
}

func handleVaultMemoryUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("entry-id") && len(unusedArgs) > 0 {
		cmd.Set("entry-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultMemoryUpdateParams{}

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

	return client.Vault.Memory.Update(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("entry-id").(string),
		params,
		options...,
	)
}

func handleVaultMemoryList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.Memory.List(ctx, cmd.Value("id").(string), options...)
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
		Title:          "vault:memory list",
		Transform:      transform,
	})
}

func handleVaultMemoryDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("entry-id") && len(unusedArgs) > 0 {
		cmd.Set("entry-id", unusedArgs[0])
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

	return client.Vault.Memory.Delete(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("entry-id").(string),
		options...,
	)
}

func handleVaultMemorySearch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultMemorySearchParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.Memory.Search(
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
		Title:          "vault:memory search",
		Transform:      transform,
	})
}
