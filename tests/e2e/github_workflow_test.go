package e2e

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGithubSourceWorkflow(t *testing.T) {
	if os.Getenv("E2E_GITHUB") != "1" {
		t.Skip("set E2E_GITHUB=1 to run GitHub E2E tests")
	}

	if os.Getenv("GITHUB_TOKEN") == "" {
		t.Skip("GITHUB_TOKEN is required for GitHub E2E tests")
	}

	bin := buildShuffle(t)
	home := t.TempDir()

	runShuffle(
		t, bin, home,
		"source", "add", "gh-cli",
		"--type", "github",
		"--owner", "cli",
		"--repo", "cli",
		"--ref", "trunk",
		"--credential-ref", "env:GITHUB_TOKEN",
	)

	listOut := runShuffle(t, bin, home, "list", "gh-cli:/cmd")
	require.Contains(t, listOut, "cmd")

	viewOut := runShuffle(t, bin, home, "view", "gh-cli:/go.mod")
	require.Contains(t, viewOut, "module github.com/cli/cli")

	downloadDir := t.TempDir()
	downloadPath := filepath.Join(downloadDir, "go.mod")

	runShuffle(t, bin, home, "download", "gh-cli:/go.mod", "--output", downloadPath)

	data, err := os.ReadFile(downloadPath)
	require.NoError(t, err)
	require.Contains(t, string(data), "module github.com/cli/cli")
}
