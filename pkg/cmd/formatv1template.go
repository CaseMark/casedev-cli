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

var formatV1TemplatesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new format template for document formatting. Templates support\nvariables using `{{variable}}` syntax and can be used for captions, signatures,\nletterheads, certificates, footers, or custom formatting needs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "content",
			Usage:    "Template content with {{variable}} placeholders",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Template name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Template type",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Template description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "styles",
			Usage:    "CSS styles for the template",
			BodyPath: "styles",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Template tags for organization",
			BodyPath: "tags",
		},
		&requestflag.Flag[[]string]{
			Name:     "variable",
			Usage:    "Template variables (auto-detected if not provided)",
			BodyPath: "variables",
		},
	},
	Action:          handleFormatV1TemplatesCreate,
	HideHelpCommand: true,
}

var formatV1TemplatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific document format template by ID. Format templates define how\ndocuments should be structured and formatted for specific legal use cases such\nas contracts, briefs, or pleadings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFormatV1TemplatesRetrieve,
	HideHelpCommand: true,
}

var formatV1TemplatesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve all format templates for the organization. Templates define reusable\ndocument formatting patterns with customizable variables for consistent legal\ndocument generation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter templates by type (e.g., contract, pleading, letter)",
			QueryPath: "type",
		},
	},
	Action:          handleFormatV1TemplatesList,
	HideHelpCommand: true,
}

func handleFormatV1TemplatesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.FormatV1TemplateNewParams{}

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
	_, err = client.Format.V1.Templates.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "format:v1:templates create", obj, format, transform)
}

func handleFormatV1TemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Format.V1.Templates.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "format:v1:templates retrieve", obj, format, transform)
}

func handleFormatV1TemplatesList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.FormatV1TemplateListParams{}

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
	_, err = client.Format.V1.Templates.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "format:v1:templates list", obj, format, transform)
}
