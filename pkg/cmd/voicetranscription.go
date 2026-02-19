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

var voiceTranscriptionCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an asynchronous transcription job for audio files. Supports two modes:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "URL of the audio file to transcribe (legacy mode, no auto-storage)",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-highlights",
			Usage:    "Automatically extract key phrases and topics",
			Default:  false,
			BodyPath: "auto_highlights",
		},
		&requestflag.Flag[string]{
			Name:     "boost-param",
			Usage:    "How much to boost custom vocabulary",
			BodyPath: "boost_param",
		},
		&requestflag.Flag[bool]{
			Name:     "content-safety-labels",
			Usage:    "Enable content moderation and safety labeling",
			Default:  false,
			BodyPath: "content_safety_labels",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "Output format for the transcript when using vault mode",
			Default:  "json",
			BodyPath: "format",
		},
		&requestflag.Flag[bool]{
			Name:     "format-text",
			Usage:    "Format text with proper capitalization",
			Default:  true,
			BodyPath: "format_text",
		},
		&requestflag.Flag[string]{
			Name:     "language-code",
			Usage:    "Language code (e.g., 'en_us', 'es', 'fr'). If not specified, language will be auto-detected",
			BodyPath: "language_code",
		},
		&requestflag.Flag[bool]{
			Name:     "language-detection",
			Usage:    "Enable automatic language detection",
			Default:  false,
			BodyPath: "language_detection",
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Usage:    "Object ID of the audio file in the vault (use with vault_id)",
			BodyPath: "object_id",
		},
		&requestflag.Flag[bool]{
			Name:     "punctuate",
			Usage:    "Add punctuation to the transcript",
			Default:  true,
			BodyPath: "punctuate",
		},
		&requestflag.Flag[bool]{
			Name:     "speaker-labels",
			Usage:    "Enable speaker identification and labeling",
			Default:  false,
			BodyPath: "speaker_labels",
		},
		&requestflag.Flag[int64]{
			Name:     "speakers-expected",
			Usage:    "Expected number of speakers (improves accuracy when known)",
			BodyPath: "speakers_expected",
		},
		&requestflag.Flag[[]string]{
			Name:     "speech-model",
			Usage:    "Priority-ordered speech models to use",
			Default:  []string{"universal-3-pro", "universal-2"},
			BodyPath: "speech_models",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Usage:    "Vault ID containing the audio file (use with object_id)",
			BodyPath: "vault_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "word-boost",
			Usage:    "Custom vocabulary words to boost (e.g., legal terms)",
			BodyPath: "word_boost",
		},
	},
	Action:          handleVoiceTranscriptionCreate,
	HideHelpCommand: true,
}

var voiceTranscriptionRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve the status and result of an audio transcription job. For vault-based\njobs, returns status and result_object_id when complete. For legacy direct URL\njobs, returns the full transcription data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVoiceTranscriptionRetrieve,
	HideHelpCommand: true,
}

var voiceTranscriptionDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a transcription job. For managed vault jobs (tr\\_\\*), also removes local\njob records and managed transcript result objects. Idempotent: returns success\nif already deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVoiceTranscriptionDelete,
	HideHelpCommand: true,
}

func handleVoiceTranscriptionCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VoiceTranscriptionNewParams{}

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
	_, err = client.Voice.Transcription.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "voice:transcription create", obj, format, transform)
}

func handleVoiceTranscriptionRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Voice.Transcription.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "voice:transcription retrieve", obj, format, transform)
}

func handleVoiceTranscriptionDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Voice.Transcription.Delete(ctx, cmd.Value("id").(string), options...)
}
