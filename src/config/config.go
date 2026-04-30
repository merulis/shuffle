package config

type Config struct {
	Sources []SourceConfig `json:"sources"`
	Default string         `json:"default,omitempty"`
}
