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

var connectorsV1InstallationsList = cli.Command{
	Name:    "list",
	Usage:   "List application installations (tenants) in this organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "application",
			QueryPath: "application",
		},
		&requestflag.Flag[string]{
			Name:      "external-tenant-id",
			QueryPath: "external_tenant_id",
		},
	},
	Action:          handleConnectorsV1InstallationsList,
	HideHelpCommand: true,
}

var connectorsV1InstallationsEnsure = cli.Command{
	Name:    "ensure",
	Usage:   "Idempotently create (or return) the installation for (application,\nexternal_tenant_id) in this organization. Send the returned installation id as\nX-Case-Installation-Id on connector requests to scope them to this tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "application",
			Usage:    `Consuming application key (e.g. "p3").`,
			Required: true,
			BodyPath: "application",
		},
		&requestflag.Flag[string]{
			Name:     "external-tenant-id",
			Usage:    "The application's own tenant identifier (e.g. a P3 organization id).",
			Required: true,
			BodyPath: "external_tenant_id",
		},
	},
	Action:          handleConnectorsV1InstallationsEnsure,
	HideHelpCommand: true,
}

func handleConnectorsV1InstallationsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1InstallationListParams{}

	return client.Connectors.V1.Installations.List(ctx, params, options...)
}

func handleConnectorsV1InstallationsEnsure(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1InstallationEnsureParams{}

	return client.Connectors.V1.Installations.Ensure(ctx, params, options...)
}
