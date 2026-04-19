package app

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/domain"
)

type ReadArtifact struct {
	storage domain.Storage
}

func NewReadArtifact(storage domain.Storage) *ReadArtifact {
	return &ReadArtifact{
		storage: storage,
	}
}

func (uc *ReadArtifact) Execute(ctx context.Context, loc domain.Locator) ([]byte, error) {
	if loc.Path == "" {
		return nil, fmt.Errorf("read artifact: empty path")
	}

	data, err := uc.storage.Read(ctx, loc)
	if err != nil {
		return nil, fmt.Errorf("User case read artidact: %w", err)
	}

	return data, nil
}
