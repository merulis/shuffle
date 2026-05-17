/*
Copyright © 2026 Vadim Kuleshvo  <merulis@yandex.ru>
*/
package main

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/cli/cmd"
	"github.com/merulis/shuffle/cli/cmd/deps"
	"github.com/merulis/shuffle/src/config"
)

func main() {
	repo := config.NewFileRepository("")
	configService := config.NewService(repo)

	root := cmd.NewCmdRoot(deps.Deps{
		ConfigService: configService,
	})

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
