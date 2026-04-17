package github

import "github.com/merulis/shuffle/src/domain"

func toArtifact(item ContentItem) domain.Artifact {
	return domain.Artifact{
		Path: item.Path,
		Name: item.Name,
		Size: item.Size,
		Type: domain.ArtifactType(item.Type),
	}
}

func toArtifactType(githubType string) domain.ArtifactType {
	switch githubType {
	case "dir":
		return domain.ArtifactDir
	case "file":
		return domain.ArtifactFile
	default:
		return domain.ArtifactFile
	}
}
