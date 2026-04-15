package domain

import "context"

type Storage interface {
	List(ctx context.Context) ([]Artifact, error)
	Read(ctx context.Context) ([]byte, error)
}
