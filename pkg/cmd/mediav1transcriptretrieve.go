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

var mediaV1TranscriptsRetrieveCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Retrieves the full transcript text for a vault transcript object or an\naudio/video source object with a completed transcription job. When object_id is\na source media object, access to that source object grants access to its\ngenerated transcript artifact.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "object-id",
			Usage:    "Object ID for either the source audio/video file or transcript artifact.",
			Required: true,
			BodyPath: "object_id",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Usage:    "Vault ID containing the source media or transcript object.",
			Required: true,
			BodyPath: "vault_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transcript",
			Usage:    "Alternative nested transcript object reference.",
			BodyPath: "transcript",
		},
	},
	Action:          handleMediaV1TranscriptsRetrieveCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"transcript": {
		&requestflag.InnerFlag[string]{
			Name:       "transcript.object-id",
			InnerField: "object_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "transcript.vault-id",
			InnerField: "vault_id",
		},
	},
})

func handleMediaV1TranscriptsRetrieveCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.MediaV1TranscriptRetrieveNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Media.V1.Transcripts.Retrieve.New(ctx, params, options...)
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
		Title:          "media:v1:transcripts:retrieve create",
		Transform:      transform,
	})
}
