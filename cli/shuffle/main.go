package main

import (
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

	ghClient := github.NewClient(token)

	if err := run(os.Args, ghClient); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string, client *github.Client) error {
	if len(args) < 2 {
		printUsage()
		return fmt.Errorf("command is required")
	}

	command := args[1]
	commandArgs := args[2:]

	switch command {
	case "ls":
		return runLs(commandArgs, client)
	case "view":
		return runView(commandArgs, client)
	case "download":
		return runDownload(commandArgs, client)
	case "help":
		printUsage()
		return nil
	default:
		printUsage()
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func printUsage() {
	fmt.Println("usage: shuffle <command> [args]")
	usageLs()
	usageView()
	usageDownload()
}
