package main

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
)

func runDownload(args []string, client *github.Client) error {
	if len(args) < 3 {
		usageView()
		return fmt.Errorf("command donwload: bad args %v", args)
	}

	ref := ""
	if len(args) >= 4 {
		ref = args[3]
	}

	dest := ""

	source := github.NewAdapter(client, args[0], args[1])
	loc := domain.NewLocator(args[2], ref)

	uc := app.NewDownloadArtifact(source)

	ctx := context.Background()

	file, err := uc.Execute(ctx, loc, dest)
	if err != nil {
		return fmt.Errorf("run command download: %w", err)
	}

	fmt.Println("saved: ", file)

	return nil
}
