package clist

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/app/cli/action"
	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service/source"
	"github.com/spf13/cobra"
)

func NewCmdList(deps cdeps.Deps) *cobra.Command {
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

func runList(deps cdeps.Deps, input string) error {
	sourceRef, err := source.ParseRef(input, ":")
	if err != nil {
		return fmt.Errorf("parse source ref: %w", err)
	}

	sourceCfg, err := deps.ConfigService.Get(sourceRef.Name)
	if err != nil {
		return fmt.Errorf("get source config: %w", err)
	}

	provider, err := deps.Factory.NewProvider(sourceCfg)
	if err != nil {
		return fmt.Errorf("create provider: %w", err)
	}

	action := action.NewListArtifacts(provider)

	items, err := action.Execute(
		context.Background(),
		entity.NewLocator(sourceRef.Path, sourceCfg.Ref),
	)
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	return nil
}
