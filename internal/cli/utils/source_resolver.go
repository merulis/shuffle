package utils

import (
	"fmt"

	"github.com/merulis/shuffle/internal/cli/deps"
	"github.com/merulis/shuffle/internal/config"
	"github.com/merulis/shuffle/internal/domain"
	"github.com/merulis/shuffle/internal/provider/factory"
)

func ResolveSource(deps deps.Deps, name string) (domain.Source, config.SourceConfig, error) {
	sourceConfig, err := deps.ConfigService.Get(name)
	if err != nil {
		return nil, config.SourceConfig{}, fmt.Errorf("get source config: %w", err)
	}

	src, err := factory.NewProvider(sourceConfig)
	if err != nil {
		return nil, config.SourceConfig{}, fmt.Errorf("create source: %w", err)
	}

	return src, sourceConfig, nil
}
