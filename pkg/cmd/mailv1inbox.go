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

var mailV1InboxesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address",
			BodyPath: "address",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			BodyPath: "displayName",
		},
	},
	Action:          handleMailV1InboxesCreate,
	HideHelpCommand: true,
}

var mailV1InboxesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesRetrieve,
	HideHelpCommand: true,
}

var mailV1InboxesList = cli.Command{
	Name:            "list",
	Usage:           "List inboxes owned by the authenticated organization.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMailV1InboxesList,
	HideHelpCommand: true,
}

var mailV1InboxesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesDelete,
	HideHelpCommand: true,
}

var mailV1InboxesGetAttachment = cli.Command{
	Name:    "get-attachment",
	Usage:   "Get attachment metadata for a message in an inbox owned by the authenticated\norganization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "attachment-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesGetAttachment,
	HideHelpCommand: true,
}

var mailV1InboxesGetMessage = cli.Command{
	Name:    "get-message",
	Usage:   "Get a message for an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesGetMessage,
	HideHelpCommand: true,
}

var mailV1InboxesListMessages = cli.Command{
	Name:    "list-messages",
	Usage:   "List messages for an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesListMessages,
	HideHelpCommand: true,
}

var mailV1InboxesReply = cli.Command{
	Name:    "reply",
	Usage:   "Reply to a message in an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesReply,
	HideHelpCommand: true,
}

var mailV1InboxesSend = cli.Command{
	Name:    "send",
	Usage:   "Send a message from an inbox owned by the authenticated organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbox-id",
			Required: true,
		},
	},
	Action:          handleMailV1InboxesSend,
	HideHelpCommand: true,
}

func handleMailV1InboxesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.MailV1InboxNewParams{}

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

	return client.Mail.V1.Inboxes.New(ctx, params, options...)
}

func handleMailV1InboxesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.Get(ctx, cmd.Value("inbox-id").(string), options...)
}

func handleMailV1InboxesList(ctx context.Context, cmd *cli.Command) error {
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

	return client.Mail.V1.Inboxes.List(ctx, options...)
}

func handleMailV1InboxesDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.Delete(ctx, cmd.Value("inbox-id").(string), options...)
}

func handleMailV1InboxesGetAttachment(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("attachment-id") && len(unusedArgs) > 0 {
		cmd.Set("attachment-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.GetAttachment(
		ctx,
		cmd.Value("inbox-id").(string),
		cmd.Value("message-id").(string),
		cmd.Value("attachment-id").(string),
		options...,
	)
}

func handleMailV1InboxesGetMessage(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.GetMessage(
		ctx,
		cmd.Value("inbox-id").(string),
		cmd.Value("message-id").(string),
		options...,
	)
}

func handleMailV1InboxesListMessages(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.ListMessages(ctx, cmd.Value("inbox-id").(string), options...)
}

func handleMailV1InboxesReply(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.Reply(
		ctx,
		cmd.Value("inbox-id").(string),
		cmd.Value("message-id").(string),
		options...,
	)
}

func handleMailV1InboxesSend(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
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

	return client.Mail.V1.Inboxes.Send(ctx, cmd.Value("inbox-id").(string), options...)
}
