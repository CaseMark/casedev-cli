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

var vaultAssembliesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Starts an asynchronous, deterministic PDF assembly. The manifest selects page\nranges from source PDFs in output order, adds an outline, an optional linked\ncontents page and optional Bates numbering, and splits the result into volumes.\nOptional compression downsamples oversized scans. Each volume becomes a new\nvault object. Source bookmarks, form fields and page labels are not carried\nover. Outputs are not indexed automatically. Completion is reported through\nvault.assembly webhooks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "source",
			Usage:    "Source PDFs, or ingested office files (spreadsheets, CSV/TSV, Word), which are assembled from the PDF version ingestion saved for them. A source can also be a PDF you made, such as a cover or notice letter, sent inline. Page ranges refer to a source by its index in this list.",
			Required: true,
			BodyPath: "sources",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "volume",
			Required: true,
			BodyPath: "volumes",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "bates",
			Usage:    "Bates numbering that continues across every volume without its own series. Text must be printable ASCII.",
			BodyPath: "bates",
		},
		&requestflag.Flag[string]{
			Name:     "client-reference",
			BodyPath: "clientReference",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "compression",
			Usage:    "Downsamples colour and grayscale scans above maxDpi and re-encodes them. Images already stored as 1-bit, JPEG 2000, indexed or CMYK keep their bytes.",
			BodyPath: "compression",
		},
		&requestflag.Flag[bool]{
			Name:     "linearize",
			Usage:    "Optimize volumes for web viewing",
			Default:  false,
			BodyPath: "linearize",
		},
		&requestflag.Flag[string]{
			Name:     "page-size",
			Usage:    "Put every page upright on a US Letter sheet: portrait pages on a portrait sheet, landscape pages on a landscape sheet, other sizes scaled to fit",
			BodyPath: "pageSize",
		},
	},
	Action:          handleVaultAssembliesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"volume": {
		&requestflag.InnerFlag[string]{
			Name:       "volume.filename",
			Usage:      "Output filename, ending in .pdf",
			InnerField: "filename",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "volume.pages",
			Usage:      "Page ranges in output order. Pages are 1-based and inclusive; repeats are allowed.",
			InnerField: "pages",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "volume.bates",
			Usage:      "This volume’s own Bates series, numbered from its own start. Same fields as the top-level bates.",
			InnerField: "bates",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "volume.contents",
			Usage:      "Adds linked contents pages at the front of the volume, built from its outline. Outline pages stay numbered within the manifest pages.",
			InnerField: "contents",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "volume.outline",
			Usage:      "Bookmarks, up to 8 levels deep",
			InnerField: "outline",
		},
		&requestflag.InnerFlag[string]{
			Name:       "volume.title",
			Usage:      "PDF document title",
			InnerField: "title",
		},
	},
	"bates": {
		&requestflag.InnerFlag[float64]{
			Name:       "bates.font-size",
			InnerField: "fontSize",
		},
		&requestflag.InnerFlag[string]{
			Name:       "bates.legend",
			Usage:      "Line printed above each label",
			InnerField: "legend",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "bates.margin-bottom",
			Usage:      "Points from the bottom edge to the label",
			InnerField: "marginBottom",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "bates.margin-side",
			Usage:      "Points from the side edge to the label",
			InnerField: "marginSide",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "bates.pad-to",
			InnerField: "padTo",
		},
		&requestflag.InnerFlag[string]{
			Name:       "bates.position",
			Usage:      `Allowed values: "bottom_right", "bottom_center", "bottom_left".`,
			InnerField: "position",
		},
		&requestflag.InnerFlag[string]{
			Name:       "bates.prefix",
			InnerField: "prefix",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "bates.shrink",
			Usage:      "Shrinks each stamped page slightly so the label sits in a clear strip below its content instead of over existing footer text",
			InnerField: "shrink",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "bates.start",
			InnerField: "start",
		},
		&requestflag.InnerFlag[string]{
			Name:       "bates.suffix",
			InnerField: "suffix",
		},
	},
	"compression": {
		&requestflag.InnerFlag[bool]{
			Name:       "compression.bitonal",
			Usage:      "Store pages that are only ink on paper as 1-bit images",
			InnerField: "bitonal",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "compression.clean-scans",
			Usage:      "Clean up scanned document pages: 300 dpi, white paper, dark ink. Photos and imaging keep their tones; gray stays gray and colour stays colour",
			InnerField: "cleanScans",
		},
		&requestflag.InnerFlag[string]{
			Name:       "compression.encoding",
			Usage:      "JPEG 2000 is smaller; JPEG is faster",
			InnerField: "encoding",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "compression.max-dpi",
			InnerField: "maxDpi",
		},
		&requestflag.InnerFlag[string]{
			Name:       "compression.quality",
			Usage:      `Allowed values: "high", "standard", "small".`,
			InnerField: "quality",
		},
	},
})

var vaultAssembliesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the status of a PDF assembly. Completed volumes include their page\ncount, size, SHA-256 and Bates range.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "assembly-id",
			Required:  true,
			PathParam: "assemblyId",
		},
	},
	Action:          handleVaultAssembliesRetrieve,
	HideHelpCommand: true,
}

func handleVaultAssembliesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.VaultAssemblyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vault.Assemblies.New(
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
		Title:          "vault:assemblies create",
		Transform:      transform,
	})
}

func handleVaultAssembliesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("assembly-id") && len(unusedArgs) > 0 {
		cmd.Set("assembly-id", unusedArgs[0])
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
	_, err = client.Vault.Assemblies.Get(
		ctx,
		cmd.Value("id").(string),
		cmd.Value("assembly-id").(string),
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
		Title:          "vault:assemblies retrieve",
		Transform:      transform,
	})
}
