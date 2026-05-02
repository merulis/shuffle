package config

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List() ([]SourceConfig, error) {
	cfg, err := s.repo.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	return cfg.Sources, nil
}

func (s *Service) Get(name string) (SourceConfig, error) {
	cfg, err := s.repo.Load()
	if err != nil {
		return SourceConfig{}, fmt.Errorf("load config: %w", err)
	}

	for _, source := range cfg.Sources {
		if source.Name == name {
			return source, nil
		}
	}

	return SourceConfig{}, fmt.Errorf("source not found")
}

func (s *Service) Add(source SourceConfig) error {
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

// TODO:
// func (s *Service) Remove(name string) error
