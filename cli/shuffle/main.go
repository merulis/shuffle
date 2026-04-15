package main

import (
	"context"
	"fmt"
	"log"

	"github.com/merulis/shuffle/src/github"
)

func main() {
	client := github.NewClient("")

	items, err := client.ListContent(
		context.Background(),
		"cli",
		"cli",
		"cmd",
		"trunk",
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}
}
