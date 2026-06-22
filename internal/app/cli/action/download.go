package action

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service"
)

type DownloadArtifact struct {
	provider service.Provider
}

func NewDownloadArtifact(provider service.Provider) *DownloadArtifact {
	return &DownloadArtifact{
		provider: provider,
	}
}

func (a *DownloadArtifact) Execute(ctx context.Context, loc entity.Locator, dest string) (string, error) {
	if loc.Path == "" {
		return "", fmt.Errorf("download artifact: empty path")
	}

	data, err := a.provider.Read(ctx, loc)
	if err != nil {
		return "", fmt.Errorf("download artifact: %w", err)
	}

	if dest == "" {
		dest = filepath.Base(loc.Path)
	}

	err = os.WriteFile(dest, data, 0o644)
	if err != nil {
		return "", fmt.Errorf("download artifact: write file: %w", err)
	}

	return dest, nil
}
