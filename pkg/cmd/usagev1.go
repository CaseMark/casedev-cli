// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	"github.com/CaseMark/casedev-go"
	"github.com/urfave/cli/v3"
)

var usageV1Retrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns customer-facing usage metrics and costs for the requested period.\nSupports summary totals and daily buckets for timestamped usage sources. Vault\nstorage is intentionally omitted from totals because it is not yet periodized\nfor arbitrary windows.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "granularity",
			Usage:     "Whether to return period totals only or include daily buckets.",
			Default:   "summary",
			QueryPath: "granularity",
		},
		&requestflag.Flag[any]{
			Name:      "period-end",
			Usage:     "Period end date. Defaults to now.",
			QueryPath: "periodEnd",
		},
		&requestflag.Flag[any]{
			Name:      "period-start",
			Usage:     "Period start date. Defaults to the start of the current calendar month.",
			QueryPath: "periodStart",
		},
	},
	Action:          handleUsageV1Retrieve,
	HideHelpCommand: true,
}

func handleUsageV1Retrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.UsageV1GetParams{}

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

	return client.Usage.V1.Get(ctx, params, options...)
}
