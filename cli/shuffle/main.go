package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env nof found, using system env")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("token is empty")
	}
	ghClient := github.NewClient(token)
	ghStorage := github.NewStorage(ghClient, "cli", "cli")
	items, err := ghStorage.List(
		context.Background(),
		domain.Locator{
			Path: "cmd",
			Ref:  "trunk",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	content, err := ghStorage.Read(
		context.Background(),
		domain.Locator{
			Path: "go.mod",
			Ref:  "trunk",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
