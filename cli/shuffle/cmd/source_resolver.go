package cmd

import (
	"fmt"

	"github.com/merulis/shuffle/src/config"
	"github.com/merulis/shuffle/src/domain"
	"github.com/merulis/shuffle/src/source"
)

func resolveSource(name string) (domain.Source, config.SourceConfig, error) {
	service := NewConfigService()

	sourceConfig, err := service.Get(name)
	if err != nil {
		return nil, config.SourceConfig{}, fmt.Errorf("get source config: %w", err)
	}

	src, err := source.New(sourceConfig)
	if err != nil {
		return nil, config.SourceConfig{}, fmt.Errorf("create source: %w", err)
	}

	return src, sourceConfig, nil
}
