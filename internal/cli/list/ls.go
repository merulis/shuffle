package list

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/source"
	"github.com/merulis/shuffle/internal/usecase"
	"github.com/spf13/cobra"
)

func NewCmdList(deps deps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <source:/path>",
		Short: "List artifacts using saved source config",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			input := args[0]
			return runList(deps, input)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	return cmd
}

func runList(deps deps.Deps, input string) error {
	sourceRef, err := source.ParseRef(input, ":")
	if err != nil {
		return err
	}

	providerSource, sourceConfig, err := deps.SourceResolver.Resolve(sourceRef)
	if err != nil {
		return err
	}

	uc := usecase.NewListArtifacts(providerSource)

	items, err := uc.Execute(
		context.Background(),
		domain.NewLocator(sourceRef.Path, sourceConfig.Ref),
	)
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	return nil
}
