package config

func validConfig() Config {
	name := "gh-cli"
	return Config{
		Sources: []SourceConfig{
			validSourceConfig(name),
		},
		Default: name,
	}
}

func validConfigWithNamedSources(names []string) Config {
	var sources []SourceConfig
	for _, name := range names {
		sources = append(sources, validSourceConfig(name))
	}
	return Config{
		Sources: sources,
		Default: names[0],
	}
}

func validSourceConfig(name string) SourceConfig {
	return SourceConfig{
		Name:  name,
		Type:  SourceTypeGithub,
		Owner: "cli",
		Repo:  "cli",
		Ref:   "trunk",
	}
}
