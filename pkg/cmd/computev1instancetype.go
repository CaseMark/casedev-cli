// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/CaseMark/casedev-cli/internal/apiquery"
	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var computeV1InstanceTypesList = cli.Command{
	Name:            "list",
	Usage:           "Retrieves all available GPU instance types with pricing, specifications, and\nregional availability. Includes T4, A10, A100, H100, and H200 GPUs powered by\nLambda Labs. Perfect for AI model training, inference workloads, and legal\ndocument OCR processing at scale.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleComputeV1InstanceTypesList,
	HideHelpCommand: true,
}

func handleComputeV1InstanceTypesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.V1.InstanceTypes.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:v1:instance-types list", obj, format, transform)
}
