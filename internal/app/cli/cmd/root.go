/*
Copyright © 2026 Vadim Kuleshov <merulis@yandex.ru>
*/
package cmd

import (
	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	cdownload "github.com/merulis/shuffle/internal/app/cli/cmd/download"
	clist "github.com/merulis/shuffle/internal/app/cli/cmd/list"
	csource "github.com/merulis/shuffle/internal/app/cli/cmd/source"
	cversion "github.com/merulis/shuffle/internal/app/cli/cmd/version"
	cview "github.com/merulis/shuffle/internal/app/cli/cmd/view"
	"github.com/spf13/cobra"
)

func NewCmdRoot(deps cdeps.Deps) *cobra.Command {
	root := &cobra.Command{
		Use:   "shuffle",
		Short: "Manage remote script and artifact sources",
		Long: `shuffle is a CLI tool for browsing, viewing, and downloading
artifacts from configured remote sources.`,
	}

	root.AddCommand(
		cversion.NewCommandVersion(),
		cview.NewCmdView(deps),
		clist.NewCmdList(deps),
		cdownload.NewCmdDownload(deps),
		csource.NewCmdSource(deps),
	)

	return root
}
