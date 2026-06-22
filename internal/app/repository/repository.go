package repository

import "github.com/merulis/shuffle/internal/app/entity"

type (
	Configuration interface {
		Load() (entity.Configuration, error)
		Save(cfg entity.Configuration) error
	}
)
