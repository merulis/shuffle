package domain

type ArtifactType string

const (
	ArtifactFile ArtifactType = "file"
	ArtifactDir  ArtifactType = "dir"
)

type Artifact struct {
	Path string
	Name string
	Size int64
	Type ArtifactType
}
