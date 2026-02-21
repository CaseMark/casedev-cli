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

var vaultObjectsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves metadata for a specific document in a vault and generates a temporary\ndownload URL. The download URL expires after 1 hour for security. This endpoint\nalso updates the file size if it wasn't previously calculated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
	},
	Action:          handleVaultObjectsRetrieve,
	HideHelpCommand: true,
}

var vaultObjectsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a document's filename, path, or metadata. Use this to rename files or\norganize them into virtual folders. The path is stored in metadata.path and can\nbe used to build folder hierarchies in your application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "filename",
			Usage:    "New filename for the document (affects display name and downloads)",
			BodyPath: "filename",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata to merge with existing metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "path",
			Usage:    "Folder path for hierarchy preservation (e.g., '/Discovery/Depositions'). Set to null or empty string to remove.",
			BodyPath: "path",
		},
	},
	Action:          handleVaultObjectsUpdate,
	HideHelpCommand: true,
}

var vaultObjectsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve all objects stored in a specific vault, including document metadata,\ningestion status, and processing statistics.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVaultObjectsList,
	HideHelpCommand: true,
}

var vaultObjectsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a document from the vault including all associated vectors,\nchunks, graph data, and the original file. This operation cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "force",
			Usage:     "Force delete a stuck document that is still in 'processing' state. Use this if a document got stuck during ingestion (e.g., OCR timeout).",
			QueryPath: "force",
		},
	},
	Action:          handleVaultObjectsDelete,
	HideHelpCommand: true,
}

var vaultObjectsCreatePresignedURL = cli.Command{
	Name:    "create-presigned-url",
	Usage:   "Generate presigned URLs for direct S3 operations (GET, PUT, DELETE, HEAD) on\nvault objects. This allows secure, time-limited access to files without proxying\nthrough the API. Essential for large document uploads/downloads in legal\nworkflows.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content-type",
			Usage:    "Content type for PUT operations (optional, defaults to object's content type)",
			BodyPath: "contentType",
		},
		&requestflag.Flag[int64]{
			Name:     "expires-in",
			Usage:    "URL expiration time in seconds (1 minute to 7 days)",
			Default:  3600,
			BodyPath: "expiresIn",
		},
		&requestflag.Flag[string]{
			Name:     "operation",
			Usage:    "The S3 operation to generate URL for",
			Default:  "GET",
			BodyPath: "operation",
		},
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "File size in bytes (optional, max 5GB for single PUT uploads). When provided for PUT operations, enforces exact file size at S3 level.",
			BodyPath: "sizeBytes",
		},
	},
	Action:          handleVaultObjectsCreatePresignedURL,
	HideHelpCommand: true,
}

var vaultObjectsGetOcrWords = cli.Command{
	Name:    "get-ocr-words",
	Usage:   "Retrieves word-level OCR bounding box data for a processed PDF document. Each\nword includes its text, normalized bounding box coordinates (0-1 range),\nconfidence score, and global word index. Use this data to highlight specific\ntext ranges in a PDF viewer based on word indices from search results.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Filter to a specific page number (1-indexed). If omitted, returns all pages.",
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "word-end",
			Usage:     "Filter to words ending at this index (inclusive). Useful for retrieving words for a specific chunk.",
			QueryPath: "wordEnd",
		},
		&requestflag.Flag[int64]{
			Name:      "word-start",
			Usage:     "Filter to words starting at this index (inclusive). Useful for retrieving words for a specific chunk.",
			QueryPath: "wordStart",
		},
	},
	Action:          handleVaultObjectsGetOcrWords,
	HideHelpCommand: true,
}

var vaultObjectsGetSummarizeJob = cli.Command{
	Name:    "get-summarize-job",
	Usage:   "Get the status of a CaseMark summary workflow job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
	},
	Action:          handleVaultObjectsGetSummarizeJob,
	HideHelpCommand: true,
}

var vaultObjectsGetText = cli.Command{
	Name:    "get-text",
	Usage:   "Retrieves the full extracted text content from a processed vault object. Returns\nthe concatenated text from all chunks, useful for document review, analysis, or\nexport. The object must have completed processing before text can be retrieved.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
	},
	Action:          handleVaultObjectsGetText,
	HideHelpCommand: true,
}

func handleVaultObjectsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
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
	_, err = client.Vault.Objects.Get(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects retrieve", obj, format, transform)
}

func handleVaultObjectsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultObjectUpdateParams{}

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
	_, err = client.Vault.Objects.Update(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects update", obj, format, transform)
}

func handleVaultObjectsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Vault.Objects.List(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects list", obj, format, transform)
}

func handleVaultObjectsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultObjectDeleteParams{}

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
	_, err = client.Vault.Objects.Delete(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects delete", obj, format, transform)
}

func handleVaultObjectsCreatePresignedURL(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultObjectNewPresignedURLParams{}

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
	_, err = client.Vault.Objects.NewPresignedURL(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects create-presigned-url", obj, format, transform)
}

func handleVaultObjectsGetOcrWords(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultObjectGetOcrWordsParams{}

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
	_, err = client.Vault.Objects.GetOcrWords(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects get-ocr-words", obj, format, transform)
}

func handleVaultObjectsGetSummarizeJob(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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
	_, err = client.Vault.Objects.GetSummarizeJob(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		cmd.Value("job-id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects get-summarize-job", obj, format, transform)
}

func handleVaultObjectsGetText(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
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
	_, err = client.Vault.Objects.GetText(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("object-id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault:objects get-text", obj, format, transform)
}
