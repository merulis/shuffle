package utils

import (
	"fmt"

	"github.com/merulis/shuffle/cli/cmd/deps"
	"github.com/merulis/shuffle/src/config"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/sourcefactory"
)

func ResolveSource(deps deps.Deps, name string) (domain.Source, config.SourceConfig, error) {
	sourceConfig, err := deps.ConfigService.Get(name)
	if err != nil {
		return nil, config.SourceConfig{}, fmt.Errorf("get source config: %w", err)
	}

	src, err := sourcefactory.New(sourceConfig)
	if err != nil {
		return nil, config.SourceConfig{}, fmt.Errorf("create source: %w", err)
	}

	return src, sourceConfig, nil
}
