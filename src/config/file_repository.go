package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultConfigDir      = ".shuffle"
	defaultConfigFileName = "config.json"
)

type FileRepository struct {
	path string
}

func NewFileRepository(path string) *FileRepository {
	if path == "" {
		path = defaultConfigPath()
	}

	return &FileRepository{
		path: path,
	}
}

func (r *FileRepository) Load() (Config, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	if len(data) == 0 {
		return Config{}, nil
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config: %w", err)
	}

	return cfg, nil
}

func (r *FileRepository) Save(cfg Config) error {
	dir := filepath.Dir(r.path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("encode Config: %w", err)
	}

	if err := os.WriteFile(r.path, data, 0644); err != nil {
		return fmt.Errorf("write config file: %w", err)
	}

	return nil
}

func defaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return defaultConfigFileName
	}

	return filepath.Join(home, defaultConfigDir, defaultConfigFileName)
}
