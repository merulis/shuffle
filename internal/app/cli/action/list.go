package action

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service"
)

type ListArtifacts struct {
	provider service.Provider
}

func NewListArtifacts(provider service.Provider) *ListArtifacts {
	return &ListArtifacts{
		provider: provider,
	}
}

func (a *ListArtifacts) Execute(ctx context.Context, loc entity.Locator) ([]entity.Artifact, error) {
	items, err := a.provider.List(ctx, loc)
	if err != nil {
		return nil, fmt.Errorf("list artifacts: %w", err)
	}

	return items, nil
}
