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

var privilegeV1Detect = cli.Command{
	Name:    "detect",
	Usage:   "Analyzes text or vault documents for legal privilege. Detects attorney-client\nprivilege, work product doctrine, common interest privilege, and litigation hold\nmaterials.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "category",
			Usage:    "Privilege categories to check. Defaults to all: attorney_client, work_product, common_interest, litigation_hold",
			BodyPath: "categories",
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Usage:    "Text content to analyze (required if document_id not provided)",
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "document-id",
			Usage:    "Vault object ID to analyze (required if content not provided)",
			BodyPath: "document_id",
		},
		&requestflag.Flag[bool]{
			Name:     "include-rationale",
			Usage:    "Include detailed rationale for each category",
			Default:  true,
			BodyPath: "include_rationale",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction",
			Usage:    "Jurisdiction for privilege rules",
			Default:  "US-Federal",
			BodyPath: "jurisdiction",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "LLM model to use for analysis",
			Default:  "casemark/casemark-core-1",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "vault-id",
			Usage:    "Vault ID (required when using document_id)",
			BodyPath: "vault_id",
		},
	},
	Action:          handlePrivilegeV1Detect,
	HideHelpCommand: true,
}

func handlePrivilegeV1Detect(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.PrivilegeV1DetectParams{}

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
	_, err = client.Privilege.V1.Detect(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "privilege:v1 detect", obj, format, transform)
}
