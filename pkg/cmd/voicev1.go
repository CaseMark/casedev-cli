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

var voiceV1ListVoices = cli.Command{
	Name:    "list-voices",
	Usage:   "Retrieve a list of available voices for text-to-speech synthesis. This endpoint\nprovides access to a comprehensive catalog of voices with various\ncharacteristics, languages, and styles suitable for legal document narration,\nclient presentations, and accessibility purposes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter by voice category",
			QueryPath: "category",
		},
		&requestflag.Flag[string]{
			Name:      "collection-id",
			Usage:     "Filter by voice collection ID",
			QueryPath: "collection_id",
		},
		&requestflag.Flag[bool]{
			Name:      "include-total-count",
			Usage:     "Whether to include total count in response",
			Default:   false,
			QueryPath: "include_total_count",
		},
		&requestflag.Flag[string]{
			Name:      "next-page-token",
			Usage:     "Token for retrieving the next page of results",
			QueryPath: "next_page_token",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of voices to return per page (max 100)",
			Default:   50,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Search term to filter voices by name or description",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Field to sort by",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "sort-direction",
			Usage:     "Sort direction",
			Default:   "asc",
			QueryPath: "sort_direction",
		},
		&requestflag.Flag[string]{
			Name:      "voice-type",
			Usage:     "Filter by voice type",
			QueryPath: "voice_type",
		},
	},
	Action:          handleVoiceV1ListVoices,
	HideHelpCommand: true,
}

func handleVoiceV1ListVoices(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.VoiceV1ListVoicesParams{}

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
	_, err = client.Voice.V1.ListVoices(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "voice:v1 list-voices", obj, format, transform)
}
