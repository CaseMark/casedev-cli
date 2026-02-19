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

var vaultCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new secure vault with dedicated S3 storage and vector search\ncapabilities. Each vault provides isolated document storage with semantic\nsearch, OCR processing, and optional GraphRAG knowledge graph features for legal\ndocument analysis and discovery.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Display name for the vault",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description of the vault's purpose",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-graph",
			Usage:    "Enable knowledge graph for entity relationship mapping. Only applies when enableIndexing is true.",
			Default:  true,
			BodyPath: "enableGraph",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-indexing",
			Usage:    "Enable vector indexing and search capabilities. Set to false for storage-only vaults.",
			Default:  true,
			BodyPath: "enableIndexing",
		},
		&requestflag.Flag[string]{
			Name:     "group-id",
			Usage:    "Assign the vault to a vault group for access control. Required when using a group-scoped API key.",
			BodyPath: "groupId",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Optional metadata to attach to the vault (e.g., { containsPHI: true } for HIPAA compliance tracking)",
			BodyPath: "metadata",
		},
	},
	Action:          handleVaultCreate,
	HideHelpCommand: true,
}

var vaultRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific vault, including storage\nconfiguration, chunking strategy, and usage statistics. Returns vault metadata,\nbucket information, and vector storage details.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVaultRetrieve,
	HideHelpCommand: true,
}

var vaultUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update vault settings including name, description, and enableGraph. Changing\nenableGraph only affects future document uploads - existing documents retain\ntheir current graph/non-graph state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "New description for the vault. Set to null to remove.",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-graph",
			Usage:    "Whether to enable GraphRAG for future document uploads",
			BodyPath: "enableGraph",
		},
		&requestflag.Flag[any]{
			Name:     "group-id",
			Usage:    "Move the vault to a different group, or set to null to remove from its current group.",
			BodyPath: "groupId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New name for the vault",
			BodyPath: "name",
		},
	},
	Action:          handleVaultUpdate,
	HideHelpCommand: true,
}

var vaultList = cli.Command{
	Name:            "list",
	Usage:           "List all vaults for the authenticated organization. Returns vault metadata\nincluding name, description, storage configuration, and usage statistics.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleVaultList,
	HideHelpCommand: true,
}

var vaultDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a vault and all its contents including documents, vectors,\ngraph data, and S3 buckets. This operation cannot be undone. For large vaults,\nuse the async=true query parameter to queue deletion in the background.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "async",
			Usage:     "If true and vault has many objects, queue deletion in background and return immediately",
			Default:   false,
			QueryPath: "async",
		},
	},
	Action:          handleVaultDelete,
	HideHelpCommand: true,
}

var vaultConfirmUpload = cli.Command{
	Name:    "confirm-upload",
	Usage:   "Confirm whether a direct-to-S3 vault upload succeeded or failed. This endpoint\nemits vault.upload.completed or vault.upload.failed events and is idempotent for\nrepeated confirmations.",
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
			Name:     "size-bytes",
			Usage:    "Uploaded file size in bytes",
			Required: true,
			BodyPath: "sizeBytes",
		},
		&requestflag.Flag[bool]{
			Name:     "success",
			Usage:    "Whether the upload succeeded",
			Required: true,
			BodyPath: "success",
		},
		&requestflag.Flag[string]{
			Name:     "etag",
			Usage:    "S3 ETag for the uploaded object (optional if client cannot access ETag header)",
			BodyPath: "etag",
		},
		&requestflag.Flag[string]{
			Name:     "error-code",
			Usage:    "Client-side error code",
			Required: true,
			BodyPath: "errorCode",
		},
		&requestflag.Flag[string]{
			Name:     "error-message",
			Usage:    "Client-side error message",
			Required: true,
			BodyPath: "errorMessage",
		},
	},
	Action:          handleVaultConfirmUpload,
	HideHelpCommand: true,
}

var vaultIngest = cli.Command{
	Name:    "ingest",
	Usage:   "Triggers ingestion workflow for a vault object to extract text, generate chunks,\nand create embeddings. For supported file types (PDF, DOCX, TXT, RTF, XML,\naudio, video), processing happens asynchronously. For unsupported types (images,\narchives, etc.), the file is marked as completed immediately without text\nextraction. GraphRAG indexing must be triggered separately via POST\n/vault/:id/graphrag/:objectId.",
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
	Action:          handleVaultIngest,
	HideHelpCommand: true,
}

var vaultSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Search across vault documents using multiple methods including hybrid vector +\ngraph search, GraphRAG global search, entity-based search, and fast similarity\nsearch. Returns relevant documents and contextual answers based on the search\nmethod.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Search query or question to find relevant documents",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filters",
			Usage:    "Filters to narrow search results to specific documents",
			BodyPath: "filters",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "Search method: 'global' for comprehensive questions, 'entity' for specific entities, 'fast' for quick similarity search, 'hybrid' for combined approach",
			Default:  "hybrid",
			BodyPath: "method",
		},
		&requestflag.Flag[int64]{
			Name:     "top-k",
			Usage:    "Maximum number of results to return",
			Default:  10,
			BodyPath: "topK",
		},
	},
	Action:          handleVaultSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.object-id",
			Usage:      "Filter to specific document(s) by object ID. Accepts a single ID or array of IDs.",
			InnerField: "object_id",
		},
	},
})

var vaultUpload = cli.Command{
	Name:    "upload",
	Usage:   "Generate a presigned URL for uploading files directly to a vault's S3 storage.\nAfter uploading to S3, confirm the upload result via POST\n/vault/:vaultId/upload/:objectId/confirm before triggering ingestion.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content-type",
			Usage:    "MIME type of the file (e.g., application/pdf, image/jpeg)",
			Required: true,
			BodyPath: "contentType",
		},
		&requestflag.Flag[string]{
			Name:     "filename",
			Usage:    "Name of the file to upload",
			Required: true,
			BodyPath: "filename",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-index",
			Usage:    "Whether to automatically process and index the file for search",
			Default:  true,
			BodyPath: "auto_index",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata to associate with the file",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "path",
			Usage:    "Optional folder path for hierarchy preservation. Allows integrations to maintain source folder structure from systems like NetDocs, Clio, or Smokeball. Example: '/Discovery/Depositions/2024'",
			BodyPath: "path",
		},
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "File size in bytes (optional, max 5GB for single PUT uploads). When provided, enforces exact file size at S3 level.",
			BodyPath: "sizeBytes",
		},
	},
	Action:          handleVaultUpload,
	HideHelpCommand: true,
}

func handleVaultCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultNewParams{}

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
	_, err = client.Vault.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault create", obj, format, transform)
}

func handleVaultRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Vault.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault retrieve", obj, format, transform)
}

func handleVaultUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultUpdateParams{}

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
	_, err = client.Vault.Update(
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
	return ShowJSON(os.Stdout, "vault update", obj, format, transform)
}

func handleVaultList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vault list", obj, format, transform)
}

func handleVaultDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultDeleteParams{}

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
	_, err = client.Vault.Delete(
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
	return ShowJSON(os.Stdout, "vault delete", obj, format, transform)
}

func handleVaultConfirmUpload(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.VaultConfirmUploadParams{}

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
	_, err = client.Vault.ConfirmUpload(
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
	return ShowJSON(os.Stdout, "vault confirm-upload", obj, format, transform)
}

func handleVaultIngest(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Vault.Ingest(
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
	return ShowJSON(os.Stdout, "vault ingest", obj, format, transform)
}

func handleVaultSearch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultSearchParams{}

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
	_, err = client.Vault.Search(
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
	return ShowJSON(os.Stdout, "vault search", obj, format, transform)
}

func handleVaultUpload(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VaultUploadParams{}

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
	_, err = client.Vault.Upload(
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
	return ShowJSON(os.Stdout, "vault upload", obj, format, transform)
}
