package csource

import (
	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/spf13/cobra"
)

func NewCmdSource(deps cdeps.Deps) *cobra.Command {
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
