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

var ocrV1Retrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve the status and results of an OCR job. Returns job progress, extracted\ntext, and metadata when processing is complete.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleOcrV1Retrieve,
	HideHelpCommand: true,
}

var ocrV1Process = requestflag.WithInnerFlags(cli.Command{
	Name:    "process",
	Usage:   "Submit a document for OCR processing to extract text, detect tables, forms, and\nother features. Supports PDFs, images, and scanned documents. Returns a job ID\nthat can be used to track processing status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "document-url",
			Usage:    "URL or S3 path to the document to process",
			Required: true,
			BodyPath: "document_url",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "URL to receive completion webhook",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[string]{
			Name:     "document-id",
			Usage:    "Optional custom document identifier",
			BodyPath: "document_id",
		},
		&requestflag.Flag[string]{
			Name:     "engine",
			Usage:    "OCR engine to use",
			Default:  "doctr",
			BodyPath: "engine",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "features",
			Usage:    "Additional processing options",
			BodyPath: "features",
		},
		&requestflag.Flag[string]{
			Name:     "result-bucket",
			Usage:    "S3 bucket to store results",
			BodyPath: "result_bucket",
		},
		&requestflag.Flag[string]{
			Name:     "result-prefix",
			Usage:    "S3 key prefix for results",
			BodyPath: "result_prefix",
		},
	},
	Action:          handleOcrV1Process,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"features": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "features.embed",
			Usage:      "Generate searchable PDF with text layer",
			InnerField: "embed",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "features.forms",
			Usage:      "Detect and extract form fields",
			InnerField: "forms",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "features.tables",
			Usage:      "Extract tables as structured data",
			InnerField: "tables",
		},
	},
})

func handleOcrV1Retrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Ocr.V1.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ocr:v1 retrieve", obj, format, transform)
}

func handleOcrV1Process(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.OcrV1ProcessParams{}

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
	_, err = client.Ocr.V1.Process(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ocr:v1 process", obj, format, transform)
}
