package pgithub

import (
	"github.com/merulis/shuffle/internal/app/entity"
)

func toArtifact(item entity.ContentItemGH) entity.Artifact {
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
