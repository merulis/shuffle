package sourcefactory

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/src/config"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/github"
)

func New(cfg config.SourceConfig) (domain.Source, error) {
	switch cfg.Type {
	case config.SourceTypeGithub:
		token := os.Getenv("GITHUB_TOKEN")

		client := github.NewClient(token)
		return github.NewAdapter(client, cfg.Owner, cfg.Repo), nil
	default:
		return nil, fmt.Errorf("unknow source type: %s", cfg.Type)
	}
}
