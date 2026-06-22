/*
Copyright © 2026 Vadim Kuleshvo  <merulis@yandex.ru>
*/
package main

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/internal/app/cli/cmd"
	cdeps "github.com/merulis/shuffle/internal/app/cli/cmd/deps"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()

	root := cmd.NewCmdRoot(cdeps.NewDeps())

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
