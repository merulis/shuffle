package source

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/cli/output"
	"github.com/spf13/cobra"
)

func NewCmdSourceList(deps deps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Add a configured source",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCmdSourceList(deps)
		},
	}

	return cmd
}

func runCmdSourceList(deps deps.Deps) error {
	sources, err := deps.ConfigService.List()
	if err != nil {
		return fmt.Errorf("list source: %w", err)
	}

	if err := output.PrintSorces(os.Stdout, sources); err != nil {
		return fmt.Errorf("print sources: %w", err)
	}

	return nil
}
