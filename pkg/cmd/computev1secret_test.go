// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestComputeV1SecretsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:v1:secrets", "create",
			"--name", "name",
			"--description", "description",
			"--env", "env",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"description: description\n" +
			"env: env\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:v1:secrets", "create",
		)
	})
}

func TestComputeV1SecretsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:v1:secrets", "list",
			"--env", "env",
		)
	})
}

func TestComputeV1SecretsDeleteGroup(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:v1:secrets", "delete-group",
			"--group", "group",
			"--env", "env",
			"--key", "key",
		)
	})
}

func TestComputeV1SecretsRetrieveGroup(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:v1:secrets", "retrieve-group",
			"--group", "group",
			"--env", "env",
		)
	})
}

func TestComputeV1SecretsUpdateGroup(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:v1:secrets", "update-group",
			"--group", "litigation-apis",
			"--secrets", "{foo: string}",
			"--env", "env",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"secrets:\n" +
			"  foo: string\n" +
			"env: env\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:v1:secrets", "update-group",
			"--group", "litigation-apis",
		)
	})
}
