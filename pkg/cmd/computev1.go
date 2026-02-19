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

var computeV1GetPricing = cli.Command{
	Name:            "get-pricing",
	Usage:           "Returns current pricing for GPU instances. Prices are fetched in real-time and\ninclude a 20% platform fee. For detailed instance types and availability, use\nGET /compute/v1/instance-types.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleComputeV1GetPricing,
	HideHelpCommand: true,
}

var computeV1GetUsage = cli.Command{
	Name:    "get-usage",
	Usage:   "Returns detailed compute usage statistics and billing information for your\norganization. Includes GPU and CPU hours, total runs, costs, and breakdowns by\nenvironment. Use optional query parameters to filter by specific year and month.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "month",
			Usage:     "Month to filter usage data (1-12, defaults to current month)",
			QueryPath: "month",
		},
		&requestflag.Flag[int64]{
			Name:      "year",
			Usage:     "Year to filter usage data (defaults to current year)",
			QueryPath: "year",
		},
	},
	Action:          handleComputeV1GetUsage,
	HideHelpCommand: true,
}

func handleComputeV1GetPricing(ctx context.Context, cmd *cli.Command) error {
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

	return client.Compute.V1.GetPricing(ctx, options...)
}

func handleComputeV1GetUsage(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.ComputeV1GetUsageParams{}

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
	_, err = client.Compute.V1.GetUsage(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1 get-usage", obj, format, transform)
}
