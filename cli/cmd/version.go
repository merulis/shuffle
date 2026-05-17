package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print shuffle version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shuffle dev")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
