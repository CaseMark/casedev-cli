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

var translateV1Detect = cli.Command{
	Name:    "detect",
	Usage:   "Detect the language of text. Returns the most likely language code and\nconfidence score. Supports batch detection for multiple texts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "q",
			Usage:    "Text to detect language for. Can be a single string or an array for batch detection.",
			Required: true,
			BodyPath: "q",
		},
	},
	Action:          handleTranslateV1Detect,
	HideHelpCommand: true,
}

var translateV1ListLanguages = cli.Command{
	Name:    "list-languages",
	Usage:   "Get the list of languages supported for translation. Optionally specify a target\nlanguage to get translated language names.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "model",
			Usage:     "Translation model to check language support for",
			QueryPath: "model",
		},
		&requestflag.Flag[string]{
			Name:      "target",
			Usage:     "Target language code for translating language names (e.g., 'es' for Spanish names)",
			QueryPath: "target",
		},
	},
	Action:          handleTranslateV1ListLanguages,
	HideHelpCommand: true,
}

var translateV1Translate = cli.Command{
	Name:    "translate",
	Usage:   "Translate text between languages using Google Cloud Translation API. Supports\n100+ languages, automatic language detection, HTML preservation, and batch\ntranslation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "q",
			Usage:    "Text to translate. Can be a single string or an array for batch translation.",
			Required: true,
			BodyPath: "q",
		},
		&requestflag.Flag[string]{
			Name:     "target",
			Usage:    "Target language code (ISO 639-1)",
			Required: true,
			BodyPath: "target",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "Format of the source text. Use 'html' to preserve HTML tags.",
			Default:  "text",
			BodyPath: "format",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Translation model. 'nmt' (Neural Machine Translation) is recommended for quality.",
			Default:  "nmt",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    "Source language code (ISO 639-1). If not specified, language is auto-detected.",
			BodyPath: "source",
		},
	},
	Action:          handleTranslateV1Translate,
	HideHelpCommand: true,
}

func handleTranslateV1Detect(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.TranslateV1DetectParams{}

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
	_, err = client.Translate.V1.Detect(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "translate:v1 detect", obj, format, transform)
}

func handleTranslateV1ListLanguages(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.TranslateV1ListLanguagesParams{}

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
	_, err = client.Translate.V1.ListLanguages(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "translate:v1 list-languages", obj, format, transform)
}

func handleTranslateV1Translate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.TranslateV1TranslateParams{}

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
	_, err = client.Translate.V1.Translate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "translate:v1 translate", obj, format, transform)
}
