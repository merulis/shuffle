package cview

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/app/cli/action"
	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service/source"
	"github.com/spf13/cobra"
)

func NewCmdView(deps cdeps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view <source:/path>",
		Short: "Preview of artifact from a source",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			input := args[0]
			return runView(deps, input)
		},
	}

	return cmd
}

func runView(deps cdeps.Deps, input string) error {
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

	loc := entity.NewLocator(sourceRef.Path, sourceCfg.Ref)

	action := action.NewReadArtifact(provider)

	ctx := context.Background()

	content, err := action.Execute(ctx, loc)
	if err != nil {
		return fmt.Errorf("run command view: %w", err)
	}

	fmt.Print(string(content))

	return nil
}
