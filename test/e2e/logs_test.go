package e2e

import (
	"testing"

	ott "github.com/opendevstack/ods-pipeline/pkg/odstasktest"
	ttr "github.com/opendevstack/ods-pipeline/pkg/tektontaskrun"
)

func TestGatherLogsTask(t *testing.T) {
	if err := ttr.RunTask(
		ttr.InNamespace(namespaceConfig.Name),
		ttr.UsingTask("ods-pipeline-tkn-logs"),
		ott.WithGitSourceWorkspace(t, "../testdata/workspaces/sample-app", namespaceConfig.Name),
		ttr.ExpectFailure(), // no pipeline run name present
	); err != nil {
		t.Fatal(err)
	}
}
