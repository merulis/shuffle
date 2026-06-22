package csource

import (
	"fmt"

	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/spf13/cobra"
)

type addOptions struct {
	name       string
	sourceType string
	owner      string
	repo       string
	ref        string
	credential string
}

func NewCmdSourceAdd(deps cdeps.Deps) *cobra.Command {
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
	cmd.Flags().StringVar(&opts.credential, "credential-ref", "", "credential ref")

	return cmd
}

func runCmdSourceAdd(deps cdeps.Deps, opts addOptions) error {
	err := deps.ConfigService.Add(entity.Source{
		Name:  opts.name,
		Type:  entity.SourceType(opts.sourceType),
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
