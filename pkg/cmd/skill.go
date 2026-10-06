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

var skillsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an org-scoped custom skill. The skill will be searchable via\n/skills/resolve alongside curated skills.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "content",
			Usage:    "Full skill content in markdown",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Skill name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "file",
			Usage:    "Optional bundled companion files installed alongside the skill as <slug>/<path> in sandbox skill directories. The complete file set may contain at most 12 MiB of decoded content.",
			BodyPath: "files",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Arbitrary metadata (author, license, etc.)",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "URL-safe slug. Auto-generated from name if omitted.",
			BodyPath: "slug",
		},
		&requestflag.Flag[string]{
			Name:     "summary",
			Usage:    "Brief description (1-2 sentences)",
			BodyPath: "summary",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags for categorization and search boosting",
			BodyPath: "tags",
		},
	},
	Action:          handleSkillsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"file": {
		&requestflag.InnerFlag[string]{
			Name:       "file.content",
			Usage:      "UTF-8 text when encoding is utf8 (max 65,536 characters), or canonical base64 when encoding is base64 (max 262,144 decoded bytes).",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.path",
			Usage:      "Relative path inside the skill directory. SKILL.md is reserved for the root skill content.",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.content-type",
			InnerField: "contentType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.encoding",
			Usage:      "How content is encoded. Omit for UTF-8 text files.",
			InnerField: "encoding",
		},
		&requestflag.InnerFlag[any]{
			Name:       "file.metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.summary",
			InnerField: "summary",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "file.tags",
			InnerField: "tags",
		},
	},
})

var skillsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an org-scoped custom skill by slug. Only provided fields are updated.\nVersion is auto-incremented.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "slug",
			Required:  true,
			PathParam: "slug",
		},
		&requestflag.Flag[string]{
			Name:     "content",
			BodyPath: "content",
		},
		&requestflag.Flag[int64]{
			Name:     "expected-version",
			Usage:    "Reject with 409 if the skill changed since this version was read.",
			BodyPath: "expectedVersion",
		},
		&requestflag.Flag[any]{
			Name:     "file",
			Usage:    "Optional replacement companion file tree, limited to 12 MiB of decoded content. Omit to leave existing bundled files unchanged; send [] to remove bundled files.",
			BodyPath: "files",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "New slug (renames the skill)",
			BodyPath: "slug",
		},
		&requestflag.Flag[*string]{
			Name:     "summary",
			BodyPath: "summary",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleSkillsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"file": {
		&requestflag.InnerFlag[string]{
			Name:                  "file.content",
			Usage:                 "UTF-8 text when encoding is utf8 (max 65,536 characters), or canonical base64 when encoding is base64 (max 262,144 decoded bytes).",
			InnerField:            "content",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "file.path",
			InnerField:            "path",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "file.content-type",
			InnerField:            "contentType",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "file.encoding",
			Usage:                 "How content is encoded. Omit for UTF-8 text files.",
			InnerField:            "encoding",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "file.metadata",
			InnerField:            "metadata",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "file.name",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "file.summary",
			InnerField:            "summary",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[[]string]{
			Name:                  "file.tags",
			InnerField:            "tags",
			OuterIsArrayOfObjects: true,
		},
	},
})

var skillsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-delete an org-scoped custom skill by slug. The skill will no longer appear\nin search results.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "slug",
			Required:  true,
			PathParam: "slug",
		},
	},
	Action:          handleSkillsDelete,
	HideHelpCommand: true,
}

var skillsCatalog = cli.Command{
	Name:    "catalog",
	Usage:   "Browse public and organization skills using one authenticated catalog. Returns\nmetadata only; skill content is loaded separately.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Number of results to skip",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Optional text search",
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "source",
			Usage:     "Optional source filter, applied after organization overrides and before pagination. Omit to browse both sources.",
			QueryPath: "source",
		},
		&requestflag.Flag[string]{
			Name:      "tag",
			Usage:     "Optional tag filter",
			QueryPath: "tag",
		},
	},
	Action:          handleSkillsCatalog,
	HideHelpCommand: true,
}

var skillsExport = cli.Command{
	Name:    "export",
	Usage:   "Export a skill as an installable filesystem tree for sandbox runtimes.\nAuthenticated org-scoped custom skills are resolved before curated skills.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "slug",
			Required:  true,
			PathParam: "slug",
		},
		&requestflag.Flag[string]{
			Name:      "target",
			Usage:     "Agent runtime skill directory convention to export for. Most callers should omit this and pass skillSlugs when creating a runtime.",
			Default:   "agents",
			QueryPath: "target",
		},
	},
	Action:          handleSkillsExport,
	HideHelpCommand: true,
}

var skillsRead = cli.Command{
	Name:    "read",
	Usage:   "Read the full content of a legal skill by its slug. Returns markdown content,\ntags, and metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "slug",
			Required:  true,
			PathParam: "slug",
		},
	},
	Action:          handleSkillsRead,
	HideHelpCommand: true,
}

var skillsResolve = cli.Command{
	Name:    "resolve",
	Usage:   "Search the Legal Skills Store using hybrid search (text + tag + semantic).\nReturns ranked results with relevance scores.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Search query string",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return (1-20)",
			Default:   10,
			QueryPath: "limit",
		},
	},
	Action:          handleSkillsResolve,
	HideHelpCommand: true,
}

func handleSkillsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.SkillNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Skills.New(ctx, params, options...)
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
		Title:          "skills create",
		Transform:      transform,
	})
}

func handleSkillsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slug") && len(unusedArgs) > 0 {
		cmd.Set("slug", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := githubcomcasemarkcasedevgo.SkillUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Skills.Update(
		ctx,
		cmd.Value("slug").(string),
		params,
		options...,
	)
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
		Title:          "skills update",
		Transform:      transform,
	})
}

func handleSkillsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slug") && len(unusedArgs) > 0 {
		cmd.Set("slug", unusedArgs[0])
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
	_, err = client.Skills.Delete(ctx, cmd.Value("slug").(string), options...)
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
		Title:          "skills delete",
		Transform:      transform,
	})
}

func handleSkillsCatalog(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.SkillCatalogParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Skills.Catalog(ctx, params, options...)
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
		Title:          "skills catalog",
		Transform:      transform,
	})
}

func handleSkillsExport(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slug") && len(unusedArgs) > 0 {
		cmd.Set("slug", unusedArgs[0])
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

	params := githubcomcasemarkcasedevgo.SkillExportParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Skills.Export(
		ctx,
		cmd.Value("slug").(string),
		params,
		options...,
	)
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
		Title:          "skills export",
		Transform:      transform,
	})
}

func handleSkillsRead(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slug") && len(unusedArgs) > 0 {
		cmd.Set("slug", unusedArgs[0])
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
	_, err = client.Skills.Read(ctx, cmd.Value("slug").(string), options...)
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
		Title:          "skills read",
		Transform:      transform,
	})
}

func handleSkillsResolve(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.SkillResolveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Skills.Resolve(ctx, params, options...)
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
		Title:          "skills resolve",
		Transform:      transform,
	})
}
