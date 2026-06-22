package service

import (
	"context"

	"github.com/merulis/shuffle/internal/app/entity"
)

type (
	Configuration interface {
		List() ([]entity.Source, error)
		Get(name string) (entity.Source, error)
		Add(source entity.Source) error
		Remove(name string) error
	}

	Provider interface {
		List(ctx context.Context, loc entity.Locator) ([]entity.Artifact, error)
		Read(ctx context.Context, loc entity.Locator) ([]byte, error)
	}
)
