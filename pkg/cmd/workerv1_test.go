// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestWorkerV1Create(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "create",
		)
	})
}

func TestWorkerV1Retrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "retrieve",
			"--id", "id",
		)
	})
}

func TestWorkerV1Delete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "delete",
			"--id", "id",
		)
	})
}

func TestWorkerV1Boot(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "boot",
			"--id", "id",
		)
	})
}

func TestWorkerV1ProxyDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "proxy-delete",
			"--id", "id",
			"--worker-path", "workerPath",
		)
	})
}

func TestWorkerV1ProxyGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "proxy-get",
			"--id", "id",
			"--worker-path", "workerPath",
		)
	})
}

func TestWorkerV1ProxyPatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "proxy-patch",
			"--id", "id",
			"--worker-path", "workerPath",
		)
	})
}

func TestWorkerV1ProxyPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "proxy-post",
			"--id", "id",
			"--worker-path", "workerPath",
		)
	})
}

func TestWorkerV1ProxyPut(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"worker:v1", "proxy-put",
			"--id", "id",
			"--worker-path", "workerPath",
		)
	})
}
