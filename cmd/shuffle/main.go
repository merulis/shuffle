/*
Copyright © 2026 Vadim Kuleshvo  <merulis@yandex.ru>
*/
package main

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/internal/cli"
	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/config"
)

func main() {
	repo := config.NewFileRepository("")
	configService := config.NewService(repo)

	root := cli.NewCmdRoot(deps.Deps{
		ConfigService: configService,
	})

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
