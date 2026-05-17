package deps

import "github.com/merulis/shuffle/src/config"

func NewConfigService() *config.Service {
	repo := config.NewFileRepository("")
	return config.NewService(repo)
}
