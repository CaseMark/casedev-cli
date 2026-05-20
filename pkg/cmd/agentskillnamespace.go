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

var agentSkillsNamespacesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a private skill namespace owned by the authenticated org and receive a\none-time bearer token used by the case-skills publisher.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "namespace-id",
			Usage:    `URL-safe slug, e.g. "curi" or "client-firm-abc". Lowercase alphanumeric with single hyphens, 2-64 chars.`,
			Required: true,
			BodyPath: "namespaceId",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
			Name:     "label",
			BodyPath: "label",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
	},
	Action:          handleAgentSkillsNamespacesCreate,
	HideHelpCommand: true,
}

var agentSkillsNamespacesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Read skill namespace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentSkillsNamespacesRetrieve,
	HideHelpCommand: true,
}

var agentSkillsNamespacesList = cli.Command{
	Name:            "list",
	Usage:           "List all active skill namespaces owned by the authenticated organization.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAgentSkillsNamespacesList,
	HideHelpCommand: true,
}

var agentSkillsNamespacesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete skill namespace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentSkillsNamespacesDelete,
	HideHelpCommand: true,
}

var agentSkillsNamespacesPublish = requestflag.WithInnerFlags(cli.Command{
	Name:    "publish",
	Usage:   "Upload a tree of skill files for the namespace. Authenticated by the namespace\nbearer token. Atomic at the version-bump level: a partial upload leaves the\nnamespace pinned to the previous version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "file",
			Required: true,
			BodyPath: "files",
		},
	},
	Action:          handleAgentSkillsNamespacesPublish,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"file": {
		&requestflag.InnerFlag[string]{
			Name:       "file.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.encoding",
			Usage:      `Allowed values: "utf8", "base64".`,
			InnerField: "encoding",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "file.content-type",
			InnerField: "contentType",
		},
	},
})

var agentSkillsNamespacesPull = cli.Command{
	Name:    "pull",
	Usage:   "Returns the active version's file manifest with short-lived presigned S3 URLs.\nSandboxes use this to materialize the tree at /workspace/.agents/skills/ before\nopencode boots.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentSkillsNamespacesPull,
	HideHelpCommand: true,
}

var agentSkillsNamespacesRotateToken = cli.Command{
	Name:    "rotate-token",
	Usage:   "Rotate skill namespace token",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentSkillsNamespacesRotateToken,
	HideHelpCommand: true,
}

func handleAgentSkillsNamespacesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.AgentSkillNamespaceNewParams{}

	return client.Agent.Skills.Namespaces.New(ctx, params, options...)
}

func handleAgentSkillsNamespacesRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	return client.Agent.Skills.Namespaces.Get(ctx, cmd.Value("id").(string), options...)
}

func handleAgentSkillsNamespacesList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Agent.Skills.Namespaces.List(ctx, options...)
}

func handleAgentSkillsNamespacesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Agent.Skills.Namespaces.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleAgentSkillsNamespacesPublish(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomcasemarkcasedevgo.AgentSkillNamespacePublishParams{}

	return client.Agent.Skills.Namespaces.Publish(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleAgentSkillsNamespacesPull(ctx context.Context, cmd *cli.Command) error {
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

	return client.Agent.Skills.Namespaces.Pull(ctx, cmd.Value("id").(string), options...)
}

func handleAgentSkillsNamespacesRotateToken(ctx context.Context, cmd *cli.Command) error {
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

	return client.Agent.Skills.Namespaces.RotateToken(ctx, cmd.Value("id").(string), options...)
}
