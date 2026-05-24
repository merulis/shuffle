/*
Copyright © 2026 Vadim Kuleshvo  <merulis@yandex.ru>
*/
package main

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/internal/cli"
	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()

	root := cli.NewCmdRoot(deps.NewDeps())

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
