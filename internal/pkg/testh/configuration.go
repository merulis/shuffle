package testh

import (
	"github.com/merulis/shuffle/internal/app/entity"
)

func ValidConfig() entity.Configuration {
	name := "gh-cli"
	return entity.Configuration{
		Sources: []entity.Source{
			ValidSourceConfig(name),
		},
		Default: name,
	}
}

func ValidConfigWithNamedSources(names []string) entity.Configuration {
	var sources []entity.Source
	for _, name := range names {
		sources = append(sources, ValidSourceConfig(name))
	}
	return entity.Configuration{
		Sources: sources,
		Default: names[0],
	}
}

func ValidSourceConfig(name string) entity.Source {
	return entity.Source{
		Name:  name,
		Type:  entity.SourceTypeGithub,
		Owner: "cli",
		Repo:  "cli",
		Ref:   "trunk",
	}
}
