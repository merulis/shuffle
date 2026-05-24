package view

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/source"
	"github.com/merulis/shuffle/internal/usecase"
	"github.com/spf13/cobra"
)

func NewCmdView(deps deps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view <owner> <repo> <path> [ref]",
		Short: "Preview of artifact from a source",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			input := args[0]
			return runView(deps, input)
		},
	}

	return cmd
}

func runView(deps deps.Deps, input string) error {
	sourceRef, err := source.ParseRef(input, ":")
	if err != nil {
		return err
	}

	providerSource, sourceConfig, err := deps.SourceResolver.Resolve(sourceRef)
	if err != nil {
		return err
	}

	loc := domain.NewLocator(sourceRef.Path, sourceConfig.Ref)

	uc := usecase.NewReadArtifact(providerSource)

	ctx := context.Background()

	content, err := uc.Execute(ctx, loc)
	if err != nil {
		return fmt.Errorf("run command view: %w", err)
	}

	fmt.Println(string(content))

	return nil
}
