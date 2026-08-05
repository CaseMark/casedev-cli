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

var vaultMultipartAbort = cli.Command{
	Name:    "abort",
	Usage:   "Abort a multipart upload and discard uploaded parts (live).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Usage:    "Vault object ID associated with the multipart upload",
			Required: true,
			BodyPath: "objectId",
		},
		&requestflag.Flag[string]{
			Name:     "upload-id",
			Usage:    "Multipart upload ID returned when the upload was initialized",
			Required: true,
			BodyPath: "uploadId",
		},
	},
	Action:          handleVaultMultipartAbort,
	HideHelpCommand: true,
}

var vaultMultipartComplete = requestflag.WithInnerFlags(cli.Command{
	Name:    "complete",
	Usage:   "Complete a multipart upload by providing the list of part numbers and ETags\n(live). Single PUT uploads are capped at 5GB; multipart default max is 16GB\n(configurable).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "File size in bytes (default max 16GB). Configure via VAULT_MULTIPART_MAX_FILE_SIZE_BYTES.",
			Required: true,
			BodyPath: "sizeBytes",
		},
		&requestflag.Flag[string]{
			Name:     "upload-id",
			Required: true,
			BodyPath: "uploadId",
		},
	},
	Action:          handleVaultMultipartComplete,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"part": {
		&requestflag.InnerFlag[string]{
			Name:       "part.etag",
			InnerField: "etag",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "part.part-number",
			InnerField: "partNumber",
		},
	},
})

var vaultMultipartGetPartURLs = requestflag.WithInnerFlags(cli.Command{
	Name:    "get-part-urls",
	Usage:   "Generate presigned URLs for individual multipart upload parts (live).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Usage:    "Vault object ID associated with the multipart upload",
			Required: true,
			BodyPath: "objectId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "part",
			Usage:    "Multipart parts that need presigned upload URLs",
			Required: true,
			BodyPath: "parts",
		},
		&requestflag.Flag[string]{
			Name:     "upload-id",
			Usage:    "Multipart upload ID returned when the upload was initialized",
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
			Usage:      "1-based multipart part number",
			InnerField: "partNumber",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "part.size-bytes",
			Usage:      "Part size in bytes (min 5MB except final part, max 5GB).",
			InnerField: "sizeBytes",
		},
	},
})

var vaultMultipartInit = cli.Command{
	Name:    "init",
	Usage:   "Initiate a multipart upload for large files (>5GB). Single PUT uploads are\ncapped at 5GB; multipart default max is 16GB (configurable). Multipart uploads\nare supported in production. Returns an uploadId and object metadata. Use part\nURLs endpoint to upload parts and complete endpoint to finalize.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "content-type",
			Usage:    "MIME type of the file",
			Required: true,
			BodyPath: "contentType",
		},
		&requestflag.Flag[string]{
			Name:     "filename",
			Usage:    "Name of the file to upload",
			Required: true,
			BodyPath: "filename",
		},
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "File size in bytes (required, default max 16GB). Configure via VAULT_MULTIPART_MAX_FILE_SIZE_BYTES.",
			Required: true,
			BodyPath: "sizeBytes",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-index",
			Usage:    "Whether to automatically process and index the file for search",
			Default:  true,
			BodyPath: "auto_index",
		},
		&requestflag.Flag[bool]{
			Name:     "is-ai-generated",
			Usage:    "Marks the file as AI-generated work product (e.g. uploaded by an agent) rather than a user-provided source document. Persisted on the object and returned by object listings so clients can distinguish provenance.",
			Default:  false,
			BodyPath: "is_ai_generated",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata to associate with the file",
			BodyPath: "metadata",
		},
		&requestflag.Flag[int64]{
			Name:     "part-size-bytes",
			Usage:    "Multipart part size in bytes (min 5MB, max 5GB). Defaults to 64MB.",
			BodyPath: "partSizeBytes",
		},
		&requestflag.Flag[string]{
			Name:     "path",
			Usage:    "Optional folder path for hierarchy preservation",
			BodyPath: "path",
		},
	},
	Action:          handleVaultMultipartInit,
	HideHelpCommand: true,
}

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

	params := githubcomcasemarkcasedevgo.VaultMultipartAbortParams{}

	return client.Vault.Multipart.Abort(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleVaultMultipartComplete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.VaultMultipartCompleteParams{}

	return client.Vault.Multipart.Complete(
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

	params := githubcomcasemarkcasedevgo.VaultMultipartGetPartURLsParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault:multipart get-part-urls",
		Transform:      transform,
	})
}

func handleVaultMultipartInit(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.VaultMultipartInitParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.Multipart.Init(
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
		Title:          "vault:multipart init",
		Transform:      transform,
	})
}
