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

var voiceBoostListExtract = cli.Command{
	Name:    "extract",
	Usage:   "Extracts a categorized word boost list from vault documents or raw text using\nLLM entity extraction. The resulting list can be passed as `word_boost` to the\ntranscription endpoint for improved accuracy.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "category",
			Usage:    "Optional filter for entity categories to extract",
			BodyPath: "categories",
		},
		&requestflag.Flag[[]string]{
			Name:     "object-id",
			Usage:    "Object IDs of documents to extract entities from (PDFs, text files)",
			BodyPath: "object_ids",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Raw text input for entity extraction (alternative to vault documents)",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Usage:    "Vault ID containing the source documents (use with object_ids)",
			BodyPath: "vault_id",
		},
	},
	Action:          handleVoiceBoostListExtract,
	HideHelpCommand: true,
}

var voiceBoostListGenerate = cli.Command{
	Name:    "generate",
	Usage:   "Generates a categorized word boost list from a completed transcription job.\nExtracts entities from the pass-1 transcript for use as `word_boost` in a second\ntranscription pass.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "transcription-job-id",
			Usage:    "Completed pass-1 transcription job ID (tr_...)",
			Required: true,
			BodyPath: "transcription_job_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "category",
			Usage:    "Optional filter for entity categories to extract",
			BodyPath: "categories",
		},
	},
	Action:          handleVoiceBoostListGenerate,
	HideHelpCommand: true,
}

func handleVoiceBoostListExtract(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VoiceBoostListExtractParams{}

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
	_, err = client.Voice.BoostList.Extract(ctx, params, options...)
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
		Title:          "voice:boost-list extract",
		Transform:      transform,
	})
}

func handleVoiceBoostListGenerate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VoiceBoostListGenerateParams{}

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
	_, err = client.Voice.BoostList.Generate(ctx, params, options...)
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
		Title:          "voice:boost-list generate",
		Transform:      transform,
	})
}
