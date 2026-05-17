package domain

import "context"

type Source interface {
	List(ctx context.Context, loc Locator) ([]Artifact, error)
	Read(ctx context.Context, loc Locator) ([]byte, error)
}
