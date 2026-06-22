package csource

import (
	"fmt"
	"os"

	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/merulis/shuffle/internal/pkg/clih"

	"github.com/spf13/cobra"
)

func NewCmdSourceList(deps cdeps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Add a configured source",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCmdSourceList(deps)
		},
	}

	return cmd
}

func runCmdSourceList(deps cdeps.Deps) error {
	sources, err := deps.ConfigService.List()
	if err != nil {
		return fmt.Errorf("list source: %w", err)
	}

	if err := clih.PrintSorces(os.Stdout, sources); err != nil {
		return fmt.Errorf("print sources: %w", err)
	}

	return nil
}
