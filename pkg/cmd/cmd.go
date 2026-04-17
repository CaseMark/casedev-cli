// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/CaseMark/casedev-cli/internal/autocomplete"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "casedev",
		Usage:     "CLI for the casedev API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "API key authentication. Use your case.dev API key (e.g., sk_case_your_api_key_here)",
				Sources: cli.EnvVars("CASEDEV_API_KEY"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "agent:v1:agents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV1AgentsCreate,
					&agentV1AgentsRetrieve,
					&agentV1AgentsUpdate,
					&agentV1AgentsList,
					&agentV1AgentsDelete,
				},
			},
			{
				Name:     "agent:v1:run",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV1RunCreate,
					&agentV1RunList,
					&agentV1RunCancel,
					&agentV1RunEvents,
					&agentV1RunExec,
					&agentV1RunGetDetails,
					&agentV1RunGetStatus,
					&agentV1RunWatch,
				},
			},
			{
				Name:     "agent:v1:execute",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV1ExecuteCreate,
				},
			},
			{
				Name:     "agent:v1:chat",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV1ChatCreate,
					&agentV1ChatDelete,
					&agentV1ChatCancel,
					&agentV1ChatReplyToQuestion,
					&agentV1ChatRespond,
					&agentV1ChatSendMessage,
					&agentV1ChatStream,
				},
			},
			{
				Name:     "agent:v1:chat:files",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV1ChatFilesList,
					&agentV1ChatFilesDownload,
				},
			},
			{
				Name:     "agent:v2:run",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV2RunCreate,
					&agentV2RunEvents,
					&agentV2RunExec,
					&agentV2RunGetDetails,
					&agentV2RunGetStatus,
				},
			},
			{
				Name:     "agent:v2:execute",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV2ExecuteCreate,
				},
			},
			{
				Name:     "agent:v2:chat",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV2ChatCreate,
					&agentV2ChatDelete,
					&agentV2ChatCancel,
					&agentV2ChatCreateStreamToken,
					&agentV2ChatReplyToQuestion,
					&agentV2ChatRespond,
					&agentV2ChatSendMessage,
					&agentV2ChatStream,
				},
			},
			{
				Name:     "agent:v2:chat:files",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentV2ChatFilesList,
					&agentV2ChatFilesDownload,
				},
			},
			{
				Name:     "system",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&systemListServices,
				},
			},
			{
				Name:     "applications:v1:deployments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&applicationsV1DeploymentsCreate,
					&applicationsV1DeploymentsRetrieve,
					&applicationsV1DeploymentsList,
					&applicationsV1DeploymentsCancel,
					&applicationsV1DeploymentsCreateFromFiles,
					&applicationsV1DeploymentsGetLogs,
					&applicationsV1DeploymentsGetStatus,
					&applicationsV1DeploymentsStream,
				},
			},
			{
				Name:     "applications:v1:projects",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&applicationsV1ProjectsCreate,
					&applicationsV1ProjectsRetrieve,
					&applicationsV1ProjectsList,
					&applicationsV1ProjectsDelete,
					&applicationsV1ProjectsCreateDeployment,
					&applicationsV1ProjectsCreateDomain,
					&applicationsV1ProjectsCreateEnv,
					&applicationsV1ProjectsDeleteDomain,
					&applicationsV1ProjectsDeleteEnv,
					&applicationsV1ProjectsGetRuntimeLogs,
					&applicationsV1ProjectsListDeployments,
					&applicationsV1ProjectsListDomains,
					&applicationsV1ProjectsListEnv,
				},
			},
			{
				Name:     "applications:v1:workflows",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&applicationsV1WorkflowsGetStatus,
				},
			},
			{
				Name:     "compute:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeV1GetPricing,
					&computeV1GetUsage,
				},
			},
			{
				Name:     "compute:v1:environments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeV1EnvironmentsCreate,
					&computeV1EnvironmentsRetrieve,
					&computeV1EnvironmentsList,
					&computeV1EnvironmentsDelete,
					&computeV1EnvironmentsSetDefault,
				},
			},
			{
				Name:     "compute:v1:instance-types",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeV1InstanceTypesList,
				},
			},
			{
				Name:     "compute:v1:instances",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeV1InstancesCreate,
					&computeV1InstancesRetrieve,
					&computeV1InstancesList,
					&computeV1InstancesDelete,
				},
			},
			{
				Name:     "compute:v1:secrets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeV1SecretsCreate,
					&computeV1SecretsList,
					&computeV1SecretsDeleteGroup,
					&computeV1SecretsRetrieveGroup,
					&computeV1SecretsUpdateGroup,
				},
			},
			{
				Name:     "database:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&databaseV1GetUsage,
				},
			},
			{
				Name:     "database:v1:projects",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&databaseV1ProjectsCreate,
					&databaseV1ProjectsRetrieve,
					&databaseV1ProjectsList,
					&databaseV1ProjectsDelete,
					&databaseV1ProjectsCreateBranch,
					&databaseV1ProjectsGetConnection,
					&databaseV1ProjectsListBranches,
				},
			},
			{
				Name:     "format:v1:templates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&formatV1TemplatesCreate,
					&formatV1TemplatesRetrieve,
					&formatV1TemplatesList,
				},
			},
			{
				Name:     "legal:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&legalV1Docket,
					&legalV1Draft,
					&legalV1Find,
					&legalV1GetCitations,
					&legalV1GetCitationsFromURL,
					&legalV1GetFullText,
					&legalV1ListCourts,
					&legalV1ListJurisdictions,
					&legalV1PatentSearch,
					&legalV1Research,
					&legalV1SecFiling,
					&legalV1Similar,
					&legalV1TrademarkSearch,
					&legalV1Verify,
				},
			},
			{
				Name:     "matters:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1Create,
					&mattersV1Retrieve,
					&mattersV1Update,
					&mattersV1List,
				},
			},
			{
				Name:     "matters:v1:agent-types",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1AgentTypesCreate,
					&mattersV1AgentTypesList,
				},
			},
			{
				Name:     "matters:v1:parties",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1PartiesCreate,
					&mattersV1PartiesRetrieve,
					&mattersV1PartiesUpdate,
					&mattersV1PartiesList,
				},
			},
			{
				Name:     "matters:v1:types",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1TypesCreate,
					&mattersV1TypesRetrieve,
					&mattersV1TypesUpdate,
					&mattersV1TypesList,
				},
			},
			{
				Name:     "matters:v1:events:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1EventsSubscriptionsCreate,
					&mattersV1EventsSubscriptionsList,
					&mattersV1EventsSubscriptionsDelete,
				},
			},
			{
				Name:     "matters:v1:log",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1LogCreate,
					&mattersV1LogList,
					&mattersV1LogExport,
				},
			},
			{
				Name:     "matters:v1:matter-parties",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1MatterPartiesCreate,
					&mattersV1MatterPartiesList,
				},
			},
			{
				Name:     "matters:v1:shares",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1SharesCreate,
					&mattersV1SharesList,
					&mattersV1SharesDelete,
				},
			},
			{
				Name:     "matters:v1:work-items",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mattersV1WorkItemsCreate,
					&mattersV1WorkItemsRetrieve,
					&mattersV1WorkItemsUpdate,
					&mattersV1WorkItemsList,
					&mattersV1WorkItemsDecide,
					&mattersV1WorkItemsListExecutions,
				},
			},
			{
				Name:     "llm",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&llmGetConfig,
				},
			},
			{
				Name:     "llm:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&llmV1CreateEmbedding,
					&llmV1ListModels,
				},
			},
			{
				Name:     "llm:v1:chat",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&llmV1ChatCreateCompletion,
				},
			},
			{
				Name:     "memory:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&memoryV1Create,
					&memoryV1Retrieve,
					&memoryV1List,
					&memoryV1Delete,
					&memoryV1DeleteAll,
					&memoryV1Search,
				},
			},
			{
				Name:     "ocr:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&ocrV1Retrieve,
					&ocrV1Download,
					&ocrV1Process,
				},
			},
			{
				Name:     "privilege:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&privilegeV1Detect,
				},
			},
			{
				Name:     "mail:v1:inboxes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mailV1InboxesCreate,
					&mailV1InboxesRetrieve,
					&mailV1InboxesList,
					&mailV1InboxesDelete,
					&mailV1InboxesGetAttachment,
					&mailV1InboxesGetMessage,
					&mailV1InboxesGetPolicy,
					&mailV1InboxesListMessages,
					&mailV1InboxesReply,
					&mailV1InboxesSend,
					&mailV1InboxesSetPolicy,
				},
			},
			{
				Name:     "skills",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&skillsCreate,
					&skillsUpdate,
					&skillsDelete,
					&skillsRead,
					&skillsResolve,
				},
			},
			{
				Name:     "skills:custom",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&skillsCustomList,
				},
			},
			{
				Name:     "search:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchV1Answer,
					&searchV1Contents,
					&searchV1Research,
					&searchV1RetrieveResearch,
					&searchV1Search,
					&searchV1Similar,
				},
			},
			{
				Name:     "translate:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&translateV1Detect,
					&translateV1ListLanguages,
					&translateV1Translate,
				},
			},
			{
				Name:     "usage:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usageV1Retrieve,
				},
			},
			{
				Name:     "usage:v1:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usageV1SubscriptionsCreate,
					&usageV1SubscriptionsUpdate,
					&usageV1SubscriptionsList,
					&usageV1SubscriptionsDelete,
					&usageV1SubscriptionsTest,
				},
			},
			{
				Name:     "vault",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultCreate,
					&vaultRetrieve,
					&vaultUpdate,
					&vaultList,
					&vaultDelete,
					&vaultConfirmUpload,
					&vaultIngest,
					&vaultSearch,
					&vaultUpload,
				},
			},
			{
				Name:     "vault:events:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultEventsSubscriptionsCreate,
					&vaultEventsSubscriptionsUpdate,
					&vaultEventsSubscriptionsList,
					&vaultEventsSubscriptionsDelete,
					&vaultEventsSubscriptionsTest,
				},
			},
			{
				Name:     "vault:graphrag",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultGraphragGetStats,
					&vaultGraphragInit,
					&vaultGraphragProcessObject,
				},
			},
			{
				Name:     "vault:groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultGroupsCreate,
					&vaultGroupsUpdate,
					&vaultGroupsList,
					&vaultGroupsDelete,
				},
			},
			{
				Name:     "vault:multipart",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultMultipartAbort,
					&vaultMultipartGetPartURLs,
				},
			},
			{
				Name:     "vault:objects",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultObjectsRetrieve,
					&vaultObjectsUpdate,
					&vaultObjectsList,
					&vaultObjectsDelete,
					&vaultObjectsCreatePresignedURL,
					&vaultObjectsDownload,
					&vaultObjectsGetOcrWords,
					&vaultObjectsGetSummarizeJob,
					&vaultObjectsGetText,
				},
			},
			{
				Name:     "vault:memory",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&vaultMemoryCreate,
					&vaultMemoryUpdate,
					&vaultMemoryList,
					&vaultMemoryDelete,
					&vaultMemorySearch,
				},
			},
			{
				Name:     "voice:streaming",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&voiceStreamingGetURL,
				},
			},
			{
				Name:     "voice:boost-list",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&voiceBoostListExtract,
					&voiceBoostListGenerate,
				},
			},
			{
				Name:     "voice:transcription",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&voiceTranscriptionCreate,
					&voiceTranscriptionRetrieve,
					&voiceTranscriptionDelete,
				},
			},
			{
				Name:     "voice:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&voiceV1ListVoices,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "casedev @manpages [-o casedev.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "casedev.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "casedev.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
