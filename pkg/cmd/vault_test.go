// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestVaultCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "create",
		"--name", "Contract Review Archive",
		"--description", "Repository for all client contract reviews and analysis",
		"--enable-graph=true",
		"--enable-indexing=true",
		"--group-id", "grp_abc123",
		"--metadata", "{containsPHI: true, hipaaCompliant: true}",
	)
}

func TestVaultRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "retrieve",
		"--id", "vault_abc123",
	)
}

func TestVaultUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "update",
		"--id", "id",
		"--description", "description",
		"--enable-graph=false",
		"--group-id", "groupId",
		"--name", "Updated Vault Name",
	)
}

func TestVaultList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "list",
	)
}

func TestVaultDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "delete",
		"--id", "id",
		"--async=true",
	)
}

func TestVaultConfirmUpload(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "confirm-upload",
		"--id", "id",
		"--object-id", "objectId",
		"--size-bytes", "1",
		"--success=true",
		"--etag", "etag",
		"--error-code", "errorCode",
		"--error-message", "errorMessage",
	)
}

func TestVaultIngest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "ingest",
		"--id", "id",
		"--object-id", "objectId",
	)
}

func TestVaultSearch(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "search",
		"--id", "id",
		"--query", "query",
		"--filters", "{object_id: string}",
		"--method", "vector",
		"--top-k", "1",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(vaultSearch)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "search",
		"--id", "id",
		"--query", "query",
		"--filters.object-id", "string",
		"--method", "vector",
		"--top-k", "1",
	)
}

func TestVaultUpload(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"vault", "upload",
		"--id", "id",
		"--content-type", "contentType",
		"--filename", "filename",
		"--auto-index=true",
		"--metadata", "{}",
		"--path", "path",
		"--size-bytes", "1",
	)
}
