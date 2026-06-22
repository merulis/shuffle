package entity

type ArtifactType string

const (
	ArtifactFile   ArtifactType = "file"
	ArtifactDir    ArtifactType = "dir"
	ArtifactBase64 ArtifactType = "base64"
)

type Artifact struct {
	Path string
	Name string
	Size int64
	Type ArtifactType
}
