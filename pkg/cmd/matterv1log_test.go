// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
)

func TestMattersV1LogCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:log", "create",
			"--id", "id",
			"--summary", "summary",
			"--details", "{foo: bar}",
			"--event-type", "event_type",
			"--work-item-id", "work_item_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"summary: summary\n" +
			"details:\n" +
			"  foo: bar\n" +
			"event_type: event_type\n" +
			"work_item_id: work_item_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:log", "create",
			"--id", "id",
		)
	})
}

func TestMattersV1LogList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:log", "list",
			"--id", "id",
			"--actor-id", "actor_id",
			"--actor-type", "actor_type",
			"--end-time", "'2019-12-27T18:11:19.117Z'",
			"--event-type", "event_type",
			"--limit", "200",
			"--offset", "0",
			"--scope", "string",
			"--start-time", "'2019-12-27T18:11:19.117Z'",
			"--work-item-id", "work_item_id",
		)
	})
}

func TestMattersV1LogExport(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"matters:v1:log", "export",
			"--id", "id",
			"--actor-id", "actor_id",
			"--actor-type", "actor_type",
			"--end-time", "'2019-12-27T18:11:19.117Z'",
			"--event-type", "event_type",
			"--format", "json",
			"--scope", "string",
			"--start-time", "'2019-12-27T18:11:19.117Z'",
			"--work-item-id", "work_item_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"actor_id: actor_id\n" +
			"actor_type: actor_type\n" +
			"end_time: '2019-12-27T18:11:19.117Z'\n" +
			"event_type: event_type\n" +
			"format: json\n" +
			"scope: string\n" +
			"start_time: '2019-12-27T18:11:19.117Z'\n" +
			"work_item_id: work_item_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"matters:v1:log", "export",
			"--id", "id",
		)
	})
}
