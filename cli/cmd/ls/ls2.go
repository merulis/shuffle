package ls

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/cli/cmd/deps"
	"github.com/merulis/shuffle/cli/cmd/utils"
	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/spf13/cobra"
)

func NewCmdList2(deps deps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list2 <source:/path>",
		Short: "List artifacts using saved source config",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList2(deps, args)
		},
	}

	return cmd
}

func runList2(deps deps.Deps, args []string) error {
	target, err := utils.ParseTarget(args[0], ":")
	if err != nil {
		return err
	}

	source, sourceConfig, err := utils.ResolveSource(deps, target.SourceName)
	if err != nil {
		return err
	}

	uc := app.NewListArtifacts(source)

	items, err := uc.Execute(
		context.Background(),
		domain.NewLocator(target.Path, sourceConfig.Ref),
	)
	if err != nil {
		return fmt.Errorf("ls2: %w", err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	return nil
}
