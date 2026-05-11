package cmd

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:     "view(show) <owner> <repo> <path> [ref]",
	Aliases: []string{"show"},
	Short:   "Preview of artifact from a source",
	Args:    cobra.RangeArgs(3, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runView(args)
	},
}

func runView(args []string) error {
	client := NewGithubCient()
	if len(args) < 3 {
		return fmt.Errorf("command view: bad args %v", args)
	}

	ref := ""
	if len(args) >= 4 {
		ref = args[3]
	}

	storage := github.NewAdapter(client, args[0], args[1])
	loc := domain.NewLocator(args[2], ref)

	uc := app.NewReadArtifact(storage)

	ctx := context.Background()

	content, err := uc.Execute(ctx, loc)
	if err != nil {
		return fmt.Errorf("run command view: %w", err)
	}

	fmt.Println(string(content))

	return nil
}
