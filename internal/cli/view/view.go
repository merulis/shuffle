package view

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/provider/github"
	"github.com/merulis/shuffle/internal/usecase"
	"github.com/spf13/cobra"
)

func NewCmdView() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "view(show) <owner> <repo> <path> [ref]",
		Aliases: []string{"show"},
		Short:   "Preview of artifact from a source",
		Args:    cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runView(args)
		},
	}

	return cmd
}

func runView(args []string) error {
	client := deps.NewGithubCient()
	if len(args) < 3 {
		return fmt.Errorf("command view: bad args %v", args)
	}

	ref := ""
	if len(args) >= 4 {
		ref = args[3]
	}

	storage := github.NewAdapter(client, args[0], args[1])
	loc := domain.NewLocator(args[2], ref)

	uc := usecase.NewReadArtifact(storage)

	ctx := context.Background()

	content, err := uc.Execute(ctx, loc)
	if err != nil {
		return fmt.Errorf("run command view: %w", err)
	}

	fmt.Println(string(content))

	return nil
}
