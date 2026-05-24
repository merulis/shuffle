package factory

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/provider/github"
	"github.com/merulis/shuffle/internal/source"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewProvider(cfg source.Source) (domain.Source, error) {
	switch cfg.Type {
	case source.SourceTypeGithub:
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			fmt.Println("GITHUB_TOKEN is not loaded")
		}

		client := github.NewClient(token)
		return github.NewAdapter(client, cfg.Owner, cfg.Repo), nil
	default:
		return nil, fmt.Errorf("unknow source type: %s", cfg.Type)
	}
}
