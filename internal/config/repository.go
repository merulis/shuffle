package config

type Repository interface {
	Load() (Config, error)
	Save(cfg Config) error
}
