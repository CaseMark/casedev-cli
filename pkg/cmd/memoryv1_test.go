// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestMemoryV1Create(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "create",
		"--message", "{content: content, role: user}",
		"--category", "category",
		"--extraction-prompt", "extraction_prompt",
		"--infer=true",
		"--metadata", "{foo: bar}",
		"--tag-1", "tag_1",
		"--tag-10", "tag_10",
		"--tag-11", "tag_11",
		"--tag-12", "tag_12",
		"--tag-2", "tag_2",
		"--tag-3", "tag_3",
		"--tag-4", "tag_4",
		"--tag-5", "tag_5",
		"--tag-6", "tag_6",
		"--tag-7", "tag_7",
		"--tag-8", "tag_8",
		"--tag-9", "tag_9",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(memoryV1Create)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "create",
		"--message.content", "content",
		"--message.role", "user",
		"--category", "category",
		"--extraction-prompt", "extraction_prompt",
		"--infer=true",
		"--metadata", "{foo: bar}",
		"--tag-1", "tag_1",
		"--tag-10", "tag_10",
		"--tag-11", "tag_11",
		"--tag-12", "tag_12",
		"--tag-2", "tag_2",
		"--tag-3", "tag_3",
		"--tag-4", "tag_4",
		"--tag-5", "tag_5",
		"--tag-6", "tag_6",
		"--tag-7", "tag_7",
		"--tag-8", "tag_8",
		"--tag-9", "tag_9",
	)
}

func TestMemoryV1Retrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "retrieve",
		"--id", "id",
	)
}

func TestMemoryV1List(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "list",
		"--category", "category",
		"--limit", "0",
		"--offset", "0",
		"--tag-1", "tag_1",
		"--tag-10", "tag_10",
		"--tag-11", "tag_11",
		"--tag-12", "tag_12",
		"--tag-2", "tag_2",
		"--tag-3", "tag_3",
		"--tag-4", "tag_4",
		"--tag-5", "tag_5",
		"--tag-6", "tag_6",
		"--tag-7", "tag_7",
		"--tag-8", "tag_8",
		"--tag-9", "tag_9",
	)
}

func TestMemoryV1Delete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "delete",
		"--id", "id",
	)
}

func TestMemoryV1DeleteAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "delete-all",
		"--tag-1", "tag_1",
		"--tag-10", "tag_10",
		"--tag-11", "tag_11",
		"--tag-12", "tag_12",
		"--tag-2", "tag_2",
		"--tag-3", "tag_3",
		"--tag-4", "tag_4",
		"--tag-5", "tag_5",
		"--tag-6", "tag_6",
		"--tag-7", "tag_7",
		"--tag-8", "tag_8",
		"--tag-9", "tag_9",
	)
}

func TestMemoryV1Search(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"memory:v1", "search",
		"--query", "query",
		"--category", "category",
		"--tag-1", "tag_1",
		"--tag-10", "tag_10",
		"--tag-11", "tag_11",
		"--tag-12", "tag_12",
		"--tag-2", "tag_2",
		"--tag-3", "tag_3",
		"--tag-4", "tag_4",
		"--tag-5", "tag_5",
		"--tag-6", "tag_6",
		"--tag-7", "tag_7",
		"--tag-8", "tag_8",
		"--tag-9", "tag_9",
		"--top-k", "1",
	)
}
