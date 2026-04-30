package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/merulis/shuffle/src/domain"
)

type DownloadArtifact struct {
	source domain.Source
}

func NewDownloadArtifact(source domain.Source) *DownloadArtifact {
	return &DownloadArtifact{
		source: source,
	}
}

func (uc *DownloadArtifact) Execute(ctx context.Context, loc domain.Locator, dest string) (string, error) {
	if loc.Path == "" {
		return "", fmt.Errorf("download artifact: empty path")
	}

	data, err := uc.source.Read(ctx, loc)
	if err != nil {
		return "", fmt.Errorf("download artifact: %w", err)
	}

	if dest == "" {
		dest = filepath.Base(loc.Path)
	}

	err = os.WriteFile(dest, data, 0644)
	if err != nil {
		return "", fmt.Errorf("download artifact: write file: %w", err)
	}

	return dest, nil
}
