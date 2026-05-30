package github

import "github.com/merulis/shuffle/internal/entity"

func toArtifact(item ContentItem) entity.Artifact {
	return entity.Artifact{
		Path: item.Path,
		Name: item.Name,
		Size: item.Size,
		Type: entity.ArtifactType(item.Type),
	}
}

func toArtifactType(githubType string) entity.ArtifactType {
	switch githubType {
	case "dir":
		return entity.ArtifactDir
	case "file":
		return entity.ArtifactFile
	default:
		return entity.ArtifactFile
	}
}
