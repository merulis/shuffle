package action

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service"
)

type ReadArtifact struct {
	provider service.Provider
}

func NewReadArtifact(provider service.Provider) *ReadArtifact {
	return &ReadArtifact{
		provider: provider,
	}
}

func (a *ReadArtifact) Execute(ctx context.Context, loc entity.Locator) ([]byte, error) {
	if loc.Path == "" {
		return nil, fmt.Errorf("read artifact: empty path")
	}

	data, err := a.provider.Read(ctx, loc)
	if err != nil {
		return nil, fmt.Errorf("read artifact: %w", err)
	}

	return data, nil
}
