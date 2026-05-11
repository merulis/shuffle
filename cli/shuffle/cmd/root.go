/*
Copyright © 2026 Vadim Kuleshov <merulis@yandex.ru>
*/
package cmd

import (
	"github.com/merulis/shuffle/cli/shuffle/cmd/version"
	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	root := &cobra.Command{
		Use:   "shuffle",
		Short: "Manage remote script and artifact sources",
		Long: `shuffle is a CLI tool for browsing, viewing, and downloading
artifacts from configured remote sources.`,
	}

	root.AddCommand(
		version.NewVersion(),
	)

	return root
}
