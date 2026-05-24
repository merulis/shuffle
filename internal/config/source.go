package config

type SourceType string

const (
	SourceTypeGithub SourceType = "github"
)

type SourceConfig struct {
	Name  string     `json:"name"`
	Type  SourceType `json:"type"`
	Owner string     `json:"owner"`
	Repo  string     `json:"repo"`
	Ref   string     `json:"ref"`
}
