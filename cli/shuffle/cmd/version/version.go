package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "dev"

func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print shuffle version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(Version)
			return nil
		},
	}
}
