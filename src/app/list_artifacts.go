package app

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/domain"
)

type ListArtifacts struct {
	source domain.Source
}

func NewListArtifacts(source domain.Source) *ListArtifacts {
	return &ListArtifacts{
		source: source,
	}
}

func (uc *ListArtifacts) Execute(ctx context.Context, loc domain.Locator) ([]domain.Artifact, error) {
	items, err := uc.source.List(ctx, loc)
	if err != nil {
		return nil, fmt.Errorf("list artifacts: %w", err)
	}

	return items, nil
}
