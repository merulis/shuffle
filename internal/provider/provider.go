package provider

import (
	"context"

	"github.com/merulis/shuffle/internal/entity"
)

type Provider interface {
	List(ctx context.Context, loc entity.Locator) ([]entity.Artifact, error)
	Read(ctx context.Context, loc entity.Locator) ([]byte, error)
}
