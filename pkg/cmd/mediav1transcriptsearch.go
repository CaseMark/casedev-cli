// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-go"
	"github.com/urfave/cli/v3"
)

var mediaV1TranscriptsSearchCreate = cli.Command{
	Name:            "create",
	Usage:           "Search transcript words",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMediaV1TranscriptsSearchCreate,
	HideHelpCommand: true,
}

func handleMediaV1TranscriptsSearchCreate(ctx context.Context, cmd *cli.Command) error {
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

	return client.Media.V1.Transcripts.Search.New(ctx, options...)
}
