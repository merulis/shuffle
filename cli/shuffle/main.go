/*
Copyright © 2026 Vadim Kuleshvo  <merulis@yandex.ru>
*/
package main

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/cli/shuffle/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
