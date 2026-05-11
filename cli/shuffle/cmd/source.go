package cmd

import "github.com/spf13/cobra"

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage configurated artifact sources",
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
