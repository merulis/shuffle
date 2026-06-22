package cdownload

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/app/cli/action"
	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service/source"
	"github.com/spf13/cobra"
)

type dowloadOptions struct {
	input  string
	output string
}

func NewCmdDownload(deps cdeps.Deps) *cobra.Command {
	opts := dowloadOptions{}

	cmd := &cobra.Command{
		Use:     "download <owner> <repo> <path> [ref]",
		Aliases: []string{"load", "pull"},
		Short:   "Download artifact from a source",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.input = args[0]
			return runDownload(deps, opts)
		},
	}

	cmd.Flags().StringVar(&opts.output, "output", "", "output path")

	return cmd
}

func runDownload(deps cdeps.Deps, opts dowloadOptions) error {
	sourceRef, err := source.ParseRef(opts.input, ":")
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

	action := action.NewDownloadArtifact(provider)

	ctx := context.Background()

	file, err := action.Execute(ctx, loc, opts.output)
	if err != nil {
		return fmt.Errorf("run command download: %w", err)
	}

	fmt.Println("saved: ", file)

	return nil
}
