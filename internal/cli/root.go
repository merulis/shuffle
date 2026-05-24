/*
Copyright © 2026 Vadim Kuleshov <merulis@yandex.ru>
*/
package cli

import (
	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/cli/download"
	"github.com/merulis/shuffle/internal/cli/list"
	"github.com/merulis/shuffle/internal/cli/source"
	"github.com/merulis/shuffle/internal/cli/version"
	"github.com/merulis/shuffle/internal/cli/view"
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
		view.NewCmdView(deps),
		list.NewCmdList(deps),
		download.NewCmdDowload(deps),
		source.NewCmdSource(deps),
	)

	return root
}
