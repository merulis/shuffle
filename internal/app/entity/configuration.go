package entity

type Configuration struct {
	Sources []Source `json:"sources"`
	Default string   `json:"default,omitempty"`
}
