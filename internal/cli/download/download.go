package download

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/source"
	"github.com/merulis/shuffle/internal/usecase"
	"github.com/spf13/cobra"
)

type dowloadOptions struct {
	input  string
	output string
}

func NewCmdDowload(deps deps.Deps) *cobra.Command {
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

func runDownload(deps deps.Deps, opts dowloadOptions) error {
	sourceRef, err := source.ParseRef(opts.input, ":")
	if err != nil {
		return err
	}

	providerSource, sourceConfig, err := deps.SourceResolver.Resolve(sourceRef)
	if err != nil {
		return err
	}

	loc := domain.NewLocator(sourceRef.Path, sourceConfig.Ref)

	uc := usecase.NewDownloadArtifact(providerSource)

	ctx := context.Background()

	file, err := uc.Execute(ctx, loc, opts.output)
	if err != nil {
		return fmt.Errorf("run command download: %w", err)
	}

	fmt.Println("saved: ", file)

	return nil
}
