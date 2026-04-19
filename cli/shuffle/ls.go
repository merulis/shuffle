package main

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
)

func runLs(args []string, client *github.Client) error {
	if len(args) < 2 {
		usageLs()
		return fmt.Errorf("command ls: bad args %v", args)
	}

	dirpath := ""
	if len(args) >= 3 {
		dirpath = args[2]
	}

	ref := ""
	if len(args) >= 4 {
		ref = args[3]
	}

	storage := github.NewStorage(client, args[0], args[1])
	loc := domain.NewLocator(dirpath, ref)

	uc := app.NewListArtifacts(storage)

	ctx := context.Background()

	items, err := uc.Execute(ctx, loc)
	if err != nil {
		return fmt.Errorf("run command ls: %w", err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	return nil
}
