package sprovider

import (
	"fmt"
	"os"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service"
	pgithub "github.com/merulis/shuffle/internal/app/service/provider/github"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewProvider(src entity.Source) (service.Provider, error) {
	switch src.Type {
	case entity.SourceTypeGithub:
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			fmt.Println("GITHUB_TOKEN is not loaded")
		}

		client := pgithub.NewClient(token)
		return pgithub.NewAdapter(client, src.Owner, src.Repo), nil
	default:
		return nil, fmt.Errorf("unknow source type: %s", src.Type)
	}
}
