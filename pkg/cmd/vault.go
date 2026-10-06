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

var vaultCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new secure vault with dedicated S3 storage and vector search\ncapabilities. Each vault provides isolated document storage with semantic search\nand OCR processing for legal document analysis and discovery.",
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
		&requestflag.Flag[string]{
			Name:     "embedding-model",
			Usage:    "Optional embedding model for this vault. Defaults to casemark/embed-v1. Determines the S3 Vectors index dimension and which model is used at both ingest and search time. The vault is locked to this model after creation — use a re-embed flow to change later. Ignored when enableIndexing is false. Note: `casemark/llama-nemotron-embed-vl-1b-v2` is a deprecated alias for `casemark/embed-v1` (retained for SDK backward compatibility); new integrations should use `casemark/embed-v1` directly.",
			Default:  "casemark/embed-v1",
			BodyPath: "embeddingModel",
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
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleVaultRetrieve,
	HideHelpCommand: true,
}

var vaultUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update vault settings including name, description, and group membership.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "New description for the vault. Set to null to remove.",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
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
	Name:    "list",
	Usage:   "List all vaults for the authenticated organization. Returns vault metadata\nincluding name, description, storage configuration, and usage statistics.\nPagination is opt-in: pass `limit` (1-200) to receive a bounded page, then\nreplay `pagination.next_cursor` as `?cursor=` while `pagination.has_more` is\ntrue. A request with neither `limit` nor `cursor` still returns every vault, and\n`pagination.limit` is null. That default will become a bounded page in a future\nrelease — paginate now to avoid the change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation cursor from `pagination.next_cursor` of the previous page. Must be replayed with the same API key scope and `query` that produced it.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-totals",
			Usage:     "When `true`, adds `totals` covering every vault matching the filters, not just this page. Scans all objects in those vaults, so request it once per filter change rather than on every page.",
			QueryPath: "include_totals",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Vaults per page (1-200). Omit to receive every vault. Supplying a cursor without a limit uses 50.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Case-insensitive substring match on the vault name.",
			QueryPath: "query",
		},
	},
	Action:          handleVaultList,
	HideHelpCommand: true,
}

var vaultDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a vault and all its contents including documents, vectors,\ngraph data, and S3 buckets. This operation cannot be undone. For large vaults,\nuse the async=true query parameter to queue deletion in the background.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Usage:   "Confirm whether a direct-to-S3 vault upload succeeded or failed. This endpoint\nemits vault.upload.completed or vault.upload.failed events and is idempotent for\nrepeated confirmations. Conditional fields: when success=true, sizeBytes is\nrequired; when success=false, errorCode and errorMessage are required. These\nrules are enforced server-side with specific 400 responses.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "object-id",
			Required:  true,
			PathParam: "objectId",
		},
		&requestflag.Flag[bool]{
			Name:     "success",
			Usage:    "Whether the upload succeeded",
			Required: true,
			BodyPath: "success",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-ingest",
			Usage:    "When true and the object was uploaded with auto_index, trigger ingestion immediately after a successful confirmation (no separate ingest call needed). The ingest outcome is reported in the `ingest` response field; an ingest failure does not fail the confirmation.",
			Default:  false,
			BodyPath: "autoIngest",
		},
		&requestflag.Flag[string]{
			Name:     "error-code",
			Usage:    "Client-side error code. Required when success=false.",
			BodyPath: "errorCode",
		},
		&requestflag.Flag[string]{
			Name:     "error-message",
			Usage:    "Client-side error message. Required when success=false.",
			BodyPath: "errorMessage",
		},
		&requestflag.Flag[string]{
			Name:     "etag",
			Usage:    "S3 ETag for the uploaded object (optional if client cannot access ETag header). Only meaningful when success=true.",
			BodyPath: "etag",
		},
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "Uploaded file size in bytes, including zero. Required when success=true and verified against S3. Empty files can be stored and transferred, but cannot be ingested.",
			BodyPath: "sizeBytes",
		},
	},
	Action:          handleVaultConfirmUpload,
	HideHelpCommand: true,
}

var vaultIngest = cli.Command{
	Name:    "ingest",
	Usage:   "Triggers ingestion workflow for a vault object to extract text, generate chunks,\nand create embeddings. For supported file types (PDF, DOCX, PPTX, XLSX, TXT,\nRTF, XML, HTML, Markdown, CSV/TSV, JSON/YAML/TOML, common source code files,\nZIP, audio, video), processing happens asynchronously. ZIP archives always\nreturn a processing response, are unpacked recursively up to 5 levels, and each\nextracted file is created as an independent vault object and ingested via the\nnormal pipeline. For unsupported types (images, etc.), the file is marked as\ncompleted immediately without text extraction.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "object-id",
			Required:  true,
			PathParam: "objectId",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "Optional callback URL for asynchronous workflow completion.",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[[]int64]{
			Name:     "page-boundary",
			Usage:    "Optional PDF pages that must begin a new chunk segment. Overlap never crosses these boundaries.",
			BodyPath: "page_boundaries",
		},
	},
	Action:          handleVaultIngest,
	HideHelpCommand: true,
}

var vaultSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Search across vault documents using hybrid vector + BM25 search (default), fast\nvector similarity search, or a simple vector fallback. Returns matching chunks\nand their source documents.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
			Usage:    "Search method: 'hybrid' for combined vector + keyword ranking (default), 'fast' for quick vector similarity search, 'vector' for a simple document listing fallback",
			Default:  "hybrid",
			BodyPath: "method",
		},
		&requestflag.Flag[int64]{
			Name:     "top-k",
			Usage:    "Maximum number of results to return. Hybrid search supports 1 to 50; other methods may support up to 100.",
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.page-range",
			Usage:      "Restrict vector-backed retrieval to chunks wholly contained in this inclusive PDF page range. Supported by vector, hybrid, and fast methods.",
			InnerField: "page_range",
		},
	},
})

var vaultUpload = cli.Command{
	Name:    "upload",
	Usage:   "Generate a presigned URL for uploading files directly to a vault's S3 storage.\nAfter uploading to S3, confirm the upload result via POST\n/vault/:vaultId/upload/:objectId/confirm before triggering ingestion.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
		&requestflag.Flag[map[string]any]{
			Name:     "file-origin",
			Usage:    "Optional client-defined provenance metadata. Returned with the object and queryable through the object-list API.",
			BodyPath: "file_origin",
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
		&requestflag.Flag[string]{
			Name:     "path",
			Usage:    "Optional folder path, excluding the filename, for hierarchy preservation. Allows integrations to maintain source folder structure from systems like NetDocs, Clio, or Smokeball. Example: '/Discovery/Depositions/2024'",
			BodyPath: "path",
		},
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "File size in bytes (optional, including zero, max 5GB for single PUT uploads). When provided, enforces exact file size at S3 level. Empty files can be stored and transferred, but cannot be ingested.",
			BodyPath: "sizeBytes",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
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

	params := githubcomcasemarkcasedevgo.VaultNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.New(ctx, params, options...)
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
		Title:          "vault create",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault retrieve",
		Transform:      transform,
	})
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

	params := githubcomcasemarkcasedevgo.VaultUpdateParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault update",
		Transform:      transform,
	})
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

	params := githubcomcasemarkcasedevgo.VaultListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.List(ctx, params, options...)
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
		Title:          "vault list",
		Transform:      transform,
	})
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

	params := githubcomcasemarkcasedevgo.VaultDeleteParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault delete",
		Transform:      transform,
	})
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

	params := githubcomcasemarkcasedevgo.VaultConfirmUploadParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault confirm-upload",
		Transform:      transform,
	})
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomcasemarkcasedevgo.VaultIngestParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.Ingest(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault ingest",
		Transform:      transform,
	})
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

	params := githubcomcasemarkcasedevgo.VaultSearchParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault search",
		Transform:      transform,
	})
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

	params := githubcomcasemarkcasedevgo.VaultUploadParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "vault upload",
		Transform:      transform,
	})
}
