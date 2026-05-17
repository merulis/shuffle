package deps

import (
	"os"

	"github.com/merulis/shuffle/src/github"
	"github.com/subosito/gotenv"
)

func NewGithubCient() *github.Client {
	_ = gotenv.Load()

	token := os.Getenv("GITHUB_TOKEN")

	return github.NewClient(token)
}
