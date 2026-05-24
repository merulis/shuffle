package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func buildShuffle(t *testing.T) string {
	t.Helper()

	tmp := t.TempDir()
	bin := filepath.Join(tmp, "shuffle")

	cmd := exec.Command("go", "build", "-o", bin, "../../cmd/shuffle")
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	require.NoError(t, err, string(out))

	return bin
}

func runShuffle(t *testing.T, bin string, home string, args ...string) string {
	t.Helper()

	cmd := exec.Command(bin, args...)
	cmd.Env = append(
		os.Environ(),
		"HOME="+home,
	)

	out, err := cmd.CombinedOutput()
	require.NoError(t, err, string(out))

	return string(out)
}

func TestSourceAddAndList(t *testing.T) {
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

	out := runShuffle(t, bin, home, "source", "list")

	require.Contains(t, out, "gh-cli")
	require.Contains(t, out, "github")
	require.Contains(t, out, "cli")
}
