package sconfiguration

import (
	"fmt"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/repository"
	"github.com/merulis/shuffle/internal/app/service"
)

type svc struct {
	repo repository.Configuration
}

func NewService(repo repository.Configuration) service.Configuration {
	return &svc{
		repo: repo,
	}
}

func (s *svc) List() ([]entity.Source, error) {
	cfg, err := s.repo.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	return cfg.Sources, nil
}

func (s *svc) Get(name string) (entity.Source, error) {
	cfg, err := s.repo.Load()
	if err != nil {
		return entity.Source{}, fmt.Errorf("load config: %w", err)
	}

	for _, source := range cfg.Sources {
		if source.Name == name {
			return source, nil
		}
	}

	return entity.Source{}, fmt.Errorf("source not found")
}

func (s *svc) Add(source entity.Source) error {
	cfg, err := s.repo.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	for _, s := range cfg.Sources {
		if s.Name == source.Name {
			return fmt.Errorf("source exists")
		}
	}

	cfg.Sources = append(cfg.Sources, source)

	err = s.repo.Save(cfg)
	if err != nil {
		return fmt.Errorf("save updated config: %w", err)
	}

	return nil
}

func (s *svc) Remove(name string) error {
	cfg, err := s.repo.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	defaultSource := cfg.Default
	sources := cfg.Sources

	var found bool
	for i, s := range sources {
		if s.Name == name {
			found = true

			if defaultSource == s.Name {
				defaultSource = ""
			}

			// remove element without preserving order
			sources[i] = sources[len(sources)-1]
			sources = sources[:len(sources)-1]

			break
		}
	}

	if !found {
		return fmt.Errorf("source not found")
	}

	err = s.repo.Save(entity.Configuration{
		Sources: sources,
		Default: defaultSource,
	})
	if err != nil {
		return fmt.Errorf("save updated config: %w", err)
	}

	return nil
}
