// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/CaseMark/casedev-cli/internal/mocktest"
	"github.com/CaseMark/casedev-cli/internal/requestflag"
)

func TestVaultAssembliesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:assemblies", "create",
			"--id", "id",
			"--source", "{objectId: objectId, rendition: searchable_pdf}",
			"--volume", "{filename: filename, pages: [{end: 1, source: 0, start: 1, bates: true}], bates: {fontSize: 6, legend: legend, marginBottom: 0, marginSide: 0, padTo: 0, position: bottom_right, prefix: prefix, shrink: true, start: 1, suffix: suffix}, contents: {depth: 1, title: title}, outline: [{page: 1, title: title, children: [{}]}], title: title}",
			"--idempotency-key", "x",
			"--bates", "{fontSize: 6, legend: legend, marginBottom: 0, marginSide: 0, padTo: 0, position: bottom_right, prefix: prefix, shrink: true, start: 1, suffix: suffix}",
			"--client-reference", "clientReference",
			"--compression", "{bitonal: true, cleanScans: true, encoding: jpeg2000, maxDpi: 72, quality: high}",
			"--linearize=true",
			"--page-size", "letter",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(vaultAssembliesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:assemblies", "create",
			"--id", "id",
			"--source.object-id", "objectId",
			"--source.rendition", "searchable_pdf",
			"--volume.filename", "filename",
			"--volume.pages", "[{end: 1, source: 0, start: 1, bates: true}]",
			"--volume.bates", "{fontSize: 6, legend: legend, marginBottom: 0, marginSide: 0, padTo: 0, position: bottom_right, prefix: prefix, shrink: true, start: 1, suffix: suffix}",
			"--volume.contents", "{depth: 1, title: title}",
			"--volume.outline", "[{page: 1, title: title, children: [{}]}]",
			"--volume.title", "title",
			"--idempotency-key", "x",
			"--bates.font-size", "6",
			"--bates.legend", "legend",
			"--bates.margin-bottom", "0",
			"--bates.margin-side", "0",
			"--bates.pad-to", "0",
			"--bates.position", "bottom_right",
			"--bates.prefix", "prefix",
			"--bates.shrink=true",
			"--bates.start", "1",
			"--bates.suffix", "suffix",
			"--client-reference", "clientReference",
			"--compression.bitonal=true",
			"--compression.clean-scans=true",
			"--compression.encoding", "jpeg2000",
			"--compression.max-dpi", "72",
			"--compression.quality", "high",
			"--linearize=true",
			"--page-size", "letter",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sources:\n" +
			"  - objectId: objectId\n" +
			"    rendition: searchable_pdf\n" +
			"volumes:\n" +
			"  - filename: filename\n" +
			"    pages:\n" +
			"      - end: 1\n" +
			"        source: 0\n" +
			"        start: 1\n" +
			"        bates: true\n" +
			"    bates:\n" +
			"      fontSize: 6\n" +
			"      legend: legend\n" +
			"      marginBottom: 0\n" +
			"      marginSide: 0\n" +
			"      padTo: 0\n" +
			"      position: bottom_right\n" +
			"      prefix: prefix\n" +
			"      shrink: true\n" +
			"      start: 1\n" +
			"      suffix: suffix\n" +
			"    contents:\n" +
			"      depth: 1\n" +
			"      title: title\n" +
			"    outline:\n" +
			"      - page: 1\n" +
			"        title: title\n" +
			"        children:\n" +
			"          - {}\n" +
			"    title: title\n" +
			"bates:\n" +
			"  fontSize: 6\n" +
			"  legend: legend\n" +
			"  marginBottom: 0\n" +
			"  marginSide: 0\n" +
			"  padTo: 0\n" +
			"  position: bottom_right\n" +
			"  prefix: prefix\n" +
			"  shrink: true\n" +
			"  start: 1\n" +
			"  suffix: suffix\n" +
			"clientReference: clientReference\n" +
			"compression:\n" +
			"  bitonal: true\n" +
			"  cleanScans: true\n" +
			"  encoding: jpeg2000\n" +
			"  maxDpi: 72\n" +
			"  quality: high\n" +
			"linearize: true\n" +
			"pageSize: letter\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"vault:assemblies", "create",
			"--id", "id",
			"--idempotency-key", "x",
		)
	})
}

func TestVaultAssembliesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"vault:assemblies", "retrieve",
			"--id", "id",
			"--assembly-id", "assemblyId",
		)
	})
}
