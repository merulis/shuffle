package app

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/domain"
)

type ReadArtifact struct {
	source domain.Source
}

func NewReadArtifact(source domain.Source) *ReadArtifact {
	return &ReadArtifact{
		source: source,
	}
}

func (uc *ReadArtifact) Execute(ctx context.Context, loc domain.Locator) ([]byte, error) {
	if loc.Path == "" {
		return nil, fmt.Errorf("read artifact: empty path")
	}

	data, err := uc.source.Read(ctx, loc)
	if err != nil {
		return nil, fmt.Errorf("read artifact: %w", err)
	}

	return data, nil
}
