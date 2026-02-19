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

var llmV1ChatCreateCompletion = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-completion",
	Usage:   "Create a completion for the provided prompt and parameters. Compatible with\nOpenAI's chat completions API. Supports 40+ models including GPT-4, Claude,\nGemini, and CaseMark legal AI models. Includes streaming support, token\ncounting, and usage tracking.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Usage:    "List of messages comprising the conversation",
			Required: true,
			BodyPath: "messages",
		},
		&requestflag.Flag[bool]{
			Name:     "casemark-show-reasoning",
			Usage:    "CaseMark-only: when true, allows reasoning fields in responses. Defaults to false (reasoning is suppressed).",
			BodyPath: "casemark_show_reasoning",
		},
		&requestflag.Flag[float64]{
			Name:     "frequency-penalty",
			Usage:    "Frequency penalty parameter",
			BodyPath: "frequency_penalty",
		},
		&requestflag.Flag[int64]{
			Name:     "max-tokens",
			Usage:    "Maximum number of tokens to generate",
			BodyPath: "max_tokens",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model to use for completion. Defaults to casemark/casemark-core-3 if not specified",
			BodyPath: "model",
		},
		&requestflag.Flag[float64]{
			Name:     "presence-penalty",
			Usage:    "Presence penalty parameter",
			BodyPath: "presence_penalty",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "Whether to stream back partial progress",
			BodyPath: "stream",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			Usage:    "Sampling temperature between 0 and 2",
			BodyPath: "temperature",
		},
		&requestflag.Flag[float64]{
			Name:     "top-p",
			Usage:    "Nucleus sampling parameter",
			BodyPath: "top_p",
		},
	},
	Action:          handleLlmV1ChatCreateCompletion,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[string]{
			Name:       "message.content",
			Usage:      "The contents of the message",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.role",
			Usage:      "The role of the message author",
			InnerField: "role",
		},
	},
})

func handleLlmV1ChatCreateCompletion(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParams{}

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
	_, err = client.Llm.V1.Chat.NewCompletion(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "llm:v1:chat create-completion", obj, format, transform)
}
