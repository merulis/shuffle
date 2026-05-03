package cmd

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls <owner> <repo> <path> [ref]",
	Short: "List artifacts in a source",
	Args:  cobra.RangeArgs(3, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLs(args)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func runLs(args []string) error {
	client := NewGithubCient()

	if len(args) < 2 {
		return fmt.Errorf("command ls: bad args %v", args)
	}

	dirpath := ""
	if len(args) >= 3 {
		dirpath = args[2]
	}

	ref := ""
	if len(args) >= 4 {
		ref = args[3]
	}

	storage := github.NewAdapter(client, args[0], args[1])
	loc := domain.NewLocator(dirpath, ref)

	uc := app.NewListArtifacts(storage)

	ctx := context.Background()

	items, err := uc.Execute(ctx, loc)
	if err != nil {
		return fmt.Errorf("run command ls: %w", err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	return nil
}
