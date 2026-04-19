package app

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/domain"
)

type ListArtifacts struct {
	storage domain.Storage
}

func NewListArtifacts(storage domain.Storage) *ListArtifacts {
	return &ListArtifacts{
		storage: storage,
	}
}

func (uc *ListArtifacts) Execute(ctx context.Context, loc domain.Locator) ([]domain.Artifact, error) {
	items, err := uc.storage.List(ctx, loc)
	if err != nil {
		return nil, fmt.Errorf("User case list artidacts: %w", err)
	}

	return items, nil
}
