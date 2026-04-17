package domain

import "context"

type Storage interface {
	List(ctx context.Context, loc Locator) ([]Artifact, error)
	Read(ctx context.Context, loc Locator) ([]byte, error)
}
