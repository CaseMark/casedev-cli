// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1WorkItemsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:work-items", "create",
			"--id", "id",
			"--title", "title",
			"--assignee-id", "assignee_id",
			"--description", "description",
			"--due-at", "'2019-12-27T18:11:19.117Z'",
			"--exit-criterion", "string",
			"--instructions", "instructions",
			"--metadata", "{foo: bar}",
			"--priority", "low",
			"--type", "task",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"title: title\n" +
			"assignee_id: assignee_id\n" +
			"description: description\n" +
			"due_at: '2019-12-27T18:11:19.117Z'\n" +
			"exit_criteria:\n" +
			"  - string\n" +
			"instructions: instructions\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"priority: low\n" +
			"type: task\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:work-items", "create",
			"--id", "id",
		)
	})
}

func TestMattersV1WorkItemsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:work-items", "retrieve",
			"--id", "id",
			"--work-item-id", "workItemId",
		)
	})
}

func TestMattersV1WorkItemsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:work-items", "update",
			"--id", "id",
			"--work-item-id", "workItemId",
			"--assignee-id", "assignee_id",
			"--completed-at", "'2019-12-27T18:11:19.117Z'",
			"--description", "description",
			"--due-at", "'2019-12-27T18:11:19.117Z'",
			"--exit-criterion", "string",
			"--instructions", "instructions",
			"--metadata", "{foo: bar}",
			"--priority", "low",
			"--started-at", "'2019-12-27T18:11:19.117Z'",
			"--status", "draft",
			"--title", "title",
			"--type", "task",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assignee_id: assignee_id\n" +
			"completed_at: '2019-12-27T18:11:19.117Z'\n" +
			"description: description\n" +
			"due_at: '2019-12-27T18:11:19.117Z'\n" +
			"exit_criteria:\n" +
			"  - string\n" +
			"instructions: instructions\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"priority: low\n" +
			"started_at: '2019-12-27T18:11:19.117Z'\n" +
			"status: draft\n" +
			"title: title\n" +
			"type: task\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:work-items", "update",
			"--id", "id",
			"--work-item-id", "workItemId",
		)
	})
}

func TestMattersV1WorkItemsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:work-items", "list",
			"--id", "id",
			"--assignee-id", "assignee_id",
			"--status", "status",
		)
	})
}

func TestMattersV1WorkItemsDecide(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:work-items", "decide",
			"--id", "id",
			"--work-item-id", "workItemId",
			"--decision", "approve",
			"--agent-type-id", "agent_type_id",
			"--metadata", "{foo: bar}",
			"--reason", "reason",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"decision: approve\n" +
			"agent_type_id: agent_type_id\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"reason: reason\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:work-items", "decide",
			"--id", "id",
			"--work-item-id", "workItemId",
		)
	})
}

func TestMattersV1WorkItemsListExecutions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:work-items", "list-executions",
			"--id", "id",
			"--work-item-id", "workItemId",
		)
	})
}
