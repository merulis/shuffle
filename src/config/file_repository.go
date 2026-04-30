package config

import (
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

func (r *FileRepository) Load() (Config,error) {
	:
}

func (r *FileRepository) Save(cfg Config) (error) {
	:
}

func defaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return defaultConfigFileName
	}

	return filepath.Join(home, defaultConfigDir, defaultConfigFileName)
}
