/*
Copyright © 2026 Vadim Kuleshov <merulis@yandex.ru>
*/
package cmd

import (
	"github.com/merulis/shuffle/cli/cmd/deps"
	"github.com/merulis/shuffle/cli/cmd/download"
	"github.com/merulis/shuffle/cli/cmd/ls"
	"github.com/merulis/shuffle/cli/cmd/source"
	"github.com/merulis/shuffle/cli/cmd/version"
	"github.com/merulis/shuffle/cli/cmd/view"
	"github.com/spf13/cobra"
)

func NewCmdRoot(deps deps.Deps) *cobra.Command {
	root := &cobra.Command{
		Use:   "shuffle",
		Short: "Manage remote script and artifact sources",
		Long: `shuffle is a CLI tool for browsing, viewing, and downloading
artifacts from configured remote sources.`,
	}

	root.AddCommand(
		version.NewCommand(),
		view.NewCmdView(),
		ls.NewCmdList(),
		ls.NewCmdList2(deps),
		download.NewCmdDowload(),
		source.NewCmdSource(deps),
	)

	return root
}
