package cmd

import (
	"fmt"
	"strings"
)

type target struct {
	SourceName string
	Path       string
}

func parseTarget(input string, sep string) (target, error) {
	parts := strings.SplitN(input, sep, 2)
	if len(parts) != 2 {
		return target{}, fmt.Errorf("ivalid target %q, expected source:/path", input)
	}

	sourceName := strings.TrimSpace(parts[0])
	path := strings.TrimSpace(parts[1])
	path = strings.TrimPrefix(path, "/")

	if sourceName == "" {
		return target{}, fmt.Errorf("source name is empty")
	}

	if path == "" {
		return target{}, fmt.Errorf("path is empty")
	}

	return target{SourceName: sourceName, Path: path}, nil
}
