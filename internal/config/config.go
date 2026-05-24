package config

import "github.com/merulis/shuffle/internal/source"

type Config struct {
	Sources []source.Source `json:"sources"`
	Default string          `json:"default,omitempty"`
}
