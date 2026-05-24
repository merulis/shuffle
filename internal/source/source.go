package source

type (
	SourceType string
)

const (
	SourceTypeGithub SourceType = "github"
)

type Source struct {
	Name          string     `json:"name"`
	Type          SourceType `json:"type"`
	Owner         string     `json:"owner"`
	Repo          string     `json:"repo"`
	Ref           string     `json:"ref"`
	CredentialRef string     `json:"credential_ref,omitempty"`
}
