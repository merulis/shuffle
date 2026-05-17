package cmd

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:     "download(load/pull) <owner> <repo> <path> [ref]",
	Aliases: []string{"load", "pull"},
	Short:   "Download artifact from a source",
	Args:    cobra.RangeArgs(3, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDownload(args)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func runDownload(args []string) error {
	client := NewGithubCient()

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

	uc := app.NewDownloadArtifact(source)

	ctx := context.Background()

	file, err := uc.Execute(ctx, loc, dest)
	if err != nil {
		return fmt.Errorf("run command download: %w", err)
	}

	fmt.Println("saved: ", file)

	return nil
}
