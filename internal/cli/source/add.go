package source

import (
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/config"
	"github.com/spf13/cobra"
)

type addOptions struct {
	name       string
	sourceType string
	owner      string
	repo       string
	ref        string
}

func NewCmdSourceAdd(deps deps.Deps) *cobra.Command {
	opts := addOptions{}

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add a configured source",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return runCmdSourceAdd(deps, opts)
		},
	}

	cmd.Flags().StringVar(&opts.sourceType, "type", "github", "source type")
	cmd.Flags().StringVar(&opts.owner, "owner", "", "source owner")
	cmd.Flags().StringVar(&opts.repo, "repo", "", "source repo")
	cmd.Flags().StringVar(&opts.ref, "ref", "", "source ref")

	return cmd
}

func runCmdSourceAdd(deps deps.Deps, opts addOptions) error {
	err := deps.ConfigService.Add(config.SourceConfig{
		Name:  opts.name,
		Type:  config.SourceType(opts.sourceType),
		Owner: opts.owner,
		Repo:  opts.repo,
		Ref:   opts.ref,
	})
	if err != nil {
		return fmt.Errorf("add source: %w", err)
	}

	fmt.Println("source added: ", opts.name)
	return nil
}
