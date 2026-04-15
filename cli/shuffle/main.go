package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	client := github.NewClient(token)

	items, err := client.GetListContent(
		context.Background(),
		"cli",
		"cli",
		"cmd",
		"trunk",
	)
	if err != nil {
		log.Fatal(err)
	}

	content, err := client.GetFileContent(
		context.Background(),
		"cli",
		"cli",
		"go.mod",
		"trunk",
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	fmt.Println(string(content))
}
