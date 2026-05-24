package source

import (
	"fmt"
	"strings"
)

type SourceRef struct {
	Name string
	Path string
}

func ParseRef(input string, sep string) (SourceRef, error) {
	parts := strings.SplitN(input, sep, 2)
	if len(parts) != 2 {
		return SourceRef{}, fmt.Errorf("invalid target %q, expected source:/path", input)
	}

	sourceName := strings.TrimSpace(parts[0])
	path := strings.TrimSpace(parts[1])
	path = strings.TrimPrefix(path, "/")

	if sourceName == "" {
		return SourceRef{}, fmt.Errorf("source name is empty")
	}

	if path == "" {
		return SourceRef{}, fmt.Errorf("path is empty")
	}

	return SourceRef{Name: sourceName, Path: path}, nil
}
