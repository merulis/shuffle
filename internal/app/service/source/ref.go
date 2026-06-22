package source

import (
	"fmt"
	"strings"

	"github.com/merulis/shuffle/internal/app/entity"
)

func ParseRef(input string, sep string) (entity.SourceRef, error) {
	parts := strings.SplitN(input, sep, 2)
	if len(parts) != 2 {
		return entity.SourceRef{}, fmt.Errorf("invalid target %q, expected source:/path", input)
	}

	sourceName := strings.TrimSpace(parts[0])
	path := strings.TrimSpace(parts[1])
	path = strings.TrimPrefix(path, "/")

	if sourceName == "" {
		return entity.SourceRef{}, fmt.Errorf("source name is empty")
	}

	if path == "" {
		return entity.SourceRef{}, fmt.Errorf("path is empty")
	}

	return entity.SourceRef{Name: sourceName, Path: path}, nil
}
