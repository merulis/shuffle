package source

import (
	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/spf13/cobra"
)

func NewCmdSource(deps deps.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Manage configurated artifact sources",
	}

	cmd.AddCommand(
		NewCmdSourceAdd(deps),
		NewCmdSourceList(deps),
	)

	return cmd
}
