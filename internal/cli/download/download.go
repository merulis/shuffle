package download

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/provider/github"
	"github.com/merulis/shuffle/internal/usecase"
	"github.com/spf13/cobra"
)

func NewCmdDowload() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download(load/pull) <owner> <repo> <path> [ref]",
		Aliases: []string{"load", "pull"},
		Short:   "Download artifact from a source",
		Args:    cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDownload(args)
		},
	}

	return cmd
}

func runDownload(args []string) error {
	client := deps.NewGithubCient()

	if len(args) < 3 {
		return fmt.Errorf("command download: bad args %v", args)
	}

	ref := ""
	if len(args) >= 4 {
		ref = args[3]
	}

	dest := ""

	source := github.NewAdapter(client, args[0], args[1])
	loc := domain.NewLocator(args[2], ref)

	uc := usecase.NewDownloadArtifact(source)

	ctx := context.Background()

	file, err := uc.Execute(ctx, loc, dest)
	if err != nil {
		return fmt.Errorf("run command download: %w", err)
	}

	fmt.Println("saved: ", file)

	return nil
}
