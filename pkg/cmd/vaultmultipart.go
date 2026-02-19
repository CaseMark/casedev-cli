// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var vaultMultipartAbort = cli.Command{
	Name:    "abort",
	Usage:   "Abort a multipart upload and discard uploaded parts (live).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
			BodyPath: "objectId",
		},
		&requestflag.Flag[string]{
			Name:     "upload-id",
			Required: true,
			BodyPath: "uploadId",
		},
	},
	Action:          handleVaultMultipartAbort,
	HideHelpCommand: true,
}

var vaultMultipartGetPartURLs = requestflag.WithInnerFlags(cli.Command{
	Name:    "get-part-urls",
	Usage:   "Generate presigned URLs for individual multipart upload parts (live).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
			BodyPath: "objectId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "part",
			Required: true,
			BodyPath: "parts",
		},
		&requestflag.Flag[string]{
			Name:     "upload-id",
			Required: true,
			BodyPath: "uploadId",
		},
	},
	Action:          handleVaultMultipartGetPartURLs,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"part": {
		&requestflag.InnerFlag[int64]{
			Name:       "part.part-number",
			InnerField: "partNumber",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "part.size-bytes",
			Usage:      "Part size in bytes (min 5MB except final part, max 5GB).",
			InnerField: "sizeBytes",
		},
	},
})

func handleVaultMultipartAbort(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultMultipartAbortParams{}

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

	return client.Vault.Multipart.Abort(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleVaultMultipartGetPartURLs(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultMultipartGetPartURLsParams{}

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
	_, err = client.Vault.Multipart.GetPartURLs(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:multipart get-part-urls", obj, format, transform)
}
