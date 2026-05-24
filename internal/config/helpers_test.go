package config

import "github.com/merulis/shuffle/internal/source"

func validConfig() Config {
	name := "gh-cli"
	return Config{
		Sources: []source.Source{
			validSourceConfig(name),
		},
		Default: name,
	}
}

func validConfigWithNamedSources(names []string) Config {
	var sources []source.Source
	for _, name := range names {
		sources = append(sources, validSourceConfig(name))
	}
	return Config{
		Sources: sources,
		Default: names[0],
	}
}

func validSourceConfig(name string) source.Source {
	return source.Source{
		Name:  name,
		Type:  source.SourceTypeGithub,
		Owner: "cli",
		Repo:  "cli",
		Ref:   "trunk",
	}
}
