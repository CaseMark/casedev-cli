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

var llmV1CreateEmbedding = cli.Command{
	Name:    "create-embedding",
	Usage:   "Create vector embeddings from text using OpenAI-compatible models. Perfect for\nsemantic search, document similarity, and building RAG systems for legal\ndocuments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "input",
			Usage:    "Text or array of texts to create embeddings for",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Embedding model to use (e.g., text-embedding-ada-002, text-embedding-3-small)",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[int64]{
			Name:     "dimensions",
			Usage:    "Number of dimensions for the embeddings (model-specific)",
			BodyPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:     "encoding-format",
			Usage:    "Format for returned embeddings",
			Default:  "float",
			BodyPath: "encoding_format",
		},
		&requestflag.Flag[string]{
			Name:     "user",
			Usage:    "Unique identifier for the end-user",
			BodyPath: "user",
		},
	},
	Action:          handleLlmV1CreateEmbedding,
	HideHelpCommand: true,
}

var llmV1ListModels = cli.Command{
	Name:            "list-models",
	Usage:           "Retrieve a list of all available language models from 40+ providers including\nOpenAI, Anthropic, Google, and Case.dev's specialized legal models. Returns\nOpenAI-compatible model metadata with pricing information.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLlmV1ListModels,
	HideHelpCommand: true,
}

func handleLlmV1CreateEmbedding(ctx context.Context, cmd *cli.Command) error {
	client := githubcomcasemarkcasedevgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomcasemarkcasedevgo.LlmV1NewEmbeddingParams{}

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
	_, err = client.Llm.V1.NewEmbedding(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "llm:v1 create-embedding", obj, format, transform)
}

func handleLlmV1ListModels(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Llm.V1.ListModels(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "llm:v1 list-models", obj, format, transform)
}
