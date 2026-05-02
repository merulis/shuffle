package cmd

import "fmt"

func usageLs() {
	fmt.Println("	shuffle ls <owner> <repo> <path> [ref]")
}

func usageView() {
	fmt.Println("	shuffle view <owner> <repo> <path> [ref]")
}

func usageDownload() {
	fmt.Println("	shuffle download <owner> <repo> <path> [ref]")
}
