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

var connectorsV1ConnectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a pending provider connection and return a one-time connect_url for the\nhosted OAuth flow. The user completes provider consent at connect_url and is\nredirected to return_url with ?connection_id=.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    `Allowed values: "box", "clio", "dropbox", "gdrive", "microsoft", "sharefile", "smokeball".`,
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			Usage:    "HTTPS URL the user is sent back to after consent.",
			Required: true,
			BodyPath: "return_url",
		},
		&requestflag.Flag[string]{
			Name:     "scope-tier",
			Usage:    "Provider-specific OAuth permission tier. Omit to use the provider's default. Use microsoft.account.read to let Microsoft choose a personal or work/school account. Existing explicit microsoft.read and microsoft.personal.read tiers remain supported. Microsoft write tiers require the separate write switch and explicit authorization for bounded exports and paired sync.",
			BodyPath: "scope_tier",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsCreate,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve one provider connection, including account identity and health.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsRetrieve,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsList = cli.Command{
	Name:    "list",
	Usage:   "List provider connections for the organization, with health status. Returns at\nmost `limit` connections (default 200, maximum 200). When `pagination.has_more`\nis true, replay `pagination.next_cursor` as `?cursor=` to fetch the following\npage. Cursors are opaque and are only valid for the exact filter set and\ninstallation/subject scope they were issued under.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation cursor from `pagination.next_cursor` of the previous page. Must be replayed with the same filters and scope that produced it.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Connections per page (1-200). Defaults to 200.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "provider",
			QueryPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pending", "healthy", "reauth_required", "revoked", "throttled".`,
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsList,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Unlink a provider account: revoke tokens at the provider and delete them.\npurge=true additionally deletes the vault documents its import links brought in.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:      "purge",
			Default:   false,
			QueryPath: "purge",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsDelete,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsBrowse = cli.Command{
	Name:    "browse",
	Usage:   "Browse the provider one level at a time. Without a site, container, or parent,\nreturns top-level resources. Pass the stable browse_ref fields returned by one\nresponse to navigate into the next level. Returns 403\nprovider_scope_insufficient when the connection scope cannot browse server-side.\nClio browsing shares request capacity with background runs; throttled responses\ninclude Retry-After when a retry deadline is known.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "container",
			Usage:     "Container id to list, or the container containing parent",
			QueryPath: "container",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "parent",
			Usage:     "Folder id to list",
			QueryPath: "parent",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Optional provider-supported search text",
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "site",
			Usage:     "Site id to list",
			QueryPath: "site",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsBrowse,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsCreateFolder = cli.Command{
	Name:    "create-folder",
	Usage:   "Create one folder at the provider directly under a browse location, then export\ninto it with POST /connectors/v1/transfer. Name the parent with the same site,\ncontainer, and parent fields browse accepts (for Clio and Smokeball, a matter\nroot is the browse_ref the matter returned). A folder of the same name already\nunder the parent is returned with 200 instead of a duplicate; a new folder\nreturns 201. Smokeball creates folders asynchronously: when the new folder is\nnot listed yet the response is 202 with code folder_creation_pending and no\nfolder. Repeat the request after a short wait to get it with 200; a repeat sent\nwhile it is still not listed asks Smokeball again and can leave two folders with\nthat name, so space repeats out. Returns 403 provider_scope_insufficient when\nthe connection was granted read access only; reconnect with write access.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Folder name. Trimmed; 1-255 characters; no '/', '\\', or control characters; not '.' or '..'. Providers may refuse further characters.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "container",
			Usage:    "Container id of the parent, from its browse_ref",
			BodyPath: "container",
		},
		&requestflag.Flag[string]{
			Name:     "parent",
			Usage:    "Folder id of the parent, from its browse_ref. Omit for a container root.",
			BodyPath: "parent",
		},
		&requestflag.Flag[string]{
			Name:     "site",
			Usage:    "Site id of the parent, from its browse_ref",
			BodyPath: "site",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsCreateFolder,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsCreateMicrosoftAdminApproval = cli.Command{
	Name:    "create-microsoft-admin-approval",
	Usage:   "Create a seven-day, replaceable IT approval link for a pending Microsoft\nconnection. Approval does not connect an account. Only delegated document read\nscopes are requested.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsCreateMicrosoftAdminApproval,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsListContents = cli.Command{
	Name:    "list-contents",
	Usage:   "List every file under one folder, at any depth, as a flat paginated list, to\npreview what importing that folder would bring in. Pass the browse_ref fields of\na folder or container. Each file carries its path relative to that folder. A\npage returns as soon as it holds files, so it may be shorter than page_size\nwhile cursor is still set; keep requesting until cursor is null. The listing\nstops after 10,000 files and sets truncated. Returns 422 contents_not_supported\nwhen the provider cannot list folder contents, and shares the provider request\nbudget with browsing and background runs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "container",
			Usage:     "Container to list, or the container holding parent",
			QueryPath: "container",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "parent",
			Usage:     "Folder to list",
			QueryPath: "parent",
		},
		&requestflag.Flag[string]{
			Name:      "site",
			Usage:     "Site of the folder",
			QueryPath: "site",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsListContents,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsResolveLink = cli.Command{
	Name:    "resolve-link",
	Usage:   "Resolve a pasted provider folder link to that folder's browse item, including\nits browse_ref, name, and path when the provider reports it cheaply. Supports\nGoogle Drive folder and open?id= links, Box folder and shared links, Dropbox\nshared links, and OneDrive/SharePoint sharing links. Returns 400\nlink_not_recognized for links of another provider or form, 422 link_not_folder\nfor file links, 422 link_not_in_account when a Dropbox or Box shared folder is\nnot in the connected account, 404 link_not_accessible when the account cannot\nsee it, and 422 link_resolve_not_supported for providers without folder links.\nShares the provider request budget with browsing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "Folder link copied from the provider, including any query string",
			Required:  true,
			QueryPath: "url",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsResolveLink,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsSearch = cli.Command{
	Name:    "search",
	Usage:   "Search the provider's files by name, and by content where the provider supports\nit; with kind=folder, search folders by name. Without a scope, searches\neverything the connection can see; pass the browse_ref fields of a folder or\ncontainer to search inside it. Folder results carry a browse_ref to open them\nwith browse. Each response states whether its folder scope covers the whole\nsubtree. Returns 422 search_not_supported when the provider has no search and\nsearch_kind_not_supported when it cannot search folders, and shares the provider\nrequest budget with browsing and background runs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Search text",
			Required:  true,
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "container",
			Usage:     "Container to search in, or the container holding parent",
			QueryPath: "container",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "kind",
			Usage:     "What to search for: files, or folders by name",
			Default:   "file",
			QueryPath: "kind",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "parent",
			Usage:     "Folder to search in",
			QueryPath: "parent",
		},
		&requestflag.Flag[string]{
			Name:      "site",
			Usage:     "Site to search in",
			QueryPath: "site",
		},
		&requestflag.Flag[string]{
			Name:       "x-case-connector-subject",
			HeaderPath: "x-case-connector-subject",
		},
	},
	Action:          handleConnectorsV1ConnectionsSearch,
	HideHelpCommand: true,
}

var connectorsV1ConnectionsUpdateAll = cli.Command{
	Name:    "update-all",
	Usage:   "Enable or disable new runs and scheduled syncs for one provider across the\nauthenticated organization or installation. This organization-wide operation\nrequires explicit confirmation. Existing credentials, links, and imported files\nare preserved; active runs are not interrupted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "confirm-organization-wide",
			Usage:    "Confirms that this change applies to every user connection in scope.",
			Required: true,
			BodyPath: "confirm_organization_wide",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Required: true,
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Required: true,
			BodyPath: "provider",
		},
	},
	Action:          handleConnectorsV1ConnectionsUpdateAll,
	HideHelpCommand: true,
}

func handleConnectorsV1ConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.New(ctx, params, options...)
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
		Title:          "connectors:v1:connections create",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.Get(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "connectors:v1:connections retrieve",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.List(ctx, params, options...)
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
		Title:          "connectors:v1:connections list",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionDeleteParams{}

	return client.Connectors.V1.Connections.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleConnectorsV1ConnectionsBrowse(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionBrowseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.Browse(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "connectors:v1:connections browse",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsCreateFolder(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewFolderParams{}

	return client.Connectors.V1.Connections.NewFolder(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleConnectorsV1ConnectionsCreateMicrosoftAdminApproval(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewMicrosoftAdminApprovalParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.NewMicrosoftAdminApproval(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "connectors:v1:connections create-microsoft-admin-approval",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsListContents(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionListContentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.ListContents(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "connectors:v1:connections list-contents",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsResolveLink(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionResolveLinkParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.ResolveLink(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "connectors:v1:connections resolve-link",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.V1.Connections.Search(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "connectors:v1:connections search",
		Transform:      transform,
	})
}

func handleConnectorsV1ConnectionsUpdateAll(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomcasemarkcasedevgo.ConnectorV1ConnectionUpdateAllParams{}

	return client.Connectors.V1.Connections.UpdateAll(ctx, params, options...)
}
