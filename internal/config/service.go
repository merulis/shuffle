package config

import (
	"fmt"

	"github.com/merulis/shuffle/internal/source"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List() ([]source.Source, error) {
	cfg, err := s.repo.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	return cfg.Sources, nil
}

func (s *Service) Get(name string) (source.Source, error) {
	cfg, err := s.repo.Load()
	if err != nil {
		return source.Source{}, fmt.Errorf("load config: %w", err)
	}

	for _, source := range cfg.Sources {
		if source.Name == name {
			return source, nil
		}
	}

	return source.Source{}, fmt.Errorf("source not found")
}

func (s *Service) Add(source source.Source) error {
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

func (s *Service) Remove(name string) error {
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

	err = s.repo.Save(Config{
		Sources: sources,
		Default: defaultSource,
	})
	if err != nil {
		return fmt.Errorf("save updated config: %w", err)
	}

	return nil
}
