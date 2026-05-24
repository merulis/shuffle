package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeRepository struct {
	cfg Config
	err error
}

func (r *fakeRepository) Load() (Config, error) {
	return r.cfg, r.err
}

func (r *fakeRepository) Save(cfg Config) error {
	r.cfg = cfg
	return r.err
}

func TestConfig_Service_List(t *testing.T) {
	tests := []struct {
		name        string
		cfg         Config
		repoErr     error
		wantErr     string
		wantSources []SourceConfig
	}{
		{
			name:        "return sources",
			cfg:         validConfig(),
			wantSources: validConfig().Sources,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRepository{tt.cfg, tt.repoErr}
			service := NewService(repo)

			sources, err := service.List()

			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantSources, sources)
		})
	}
}

func TestConfig_Service_Get(t *testing.T) {
	tests := []struct {
		name       string
		cfg        Config
		repoErr    error
		wantErr    string
		wantName   string
		wantSource SourceConfig
	}{
		{
			name: "found source",
			cfg: validConfigWithNamedSources(
				[]string{"gh", "nexus", "gl"},
			),
			wantName:   "nexus",
			wantSource: validSourceConfig("nexus"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRepository{tt.cfg, tt.repoErr}
			service := NewService(repo)

			source, err := service.Get(tt.wantName)

			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantName, source.Name)
			assert.Equal(t, tt.wantSource, source)
		})
	}
}

func TestConfig_Service_Add(t *testing.T) {
	tests := []struct {
		name        string
		cfg         Config
		repoErr     error
		wantErr     string
		wantSource  SourceConfig
		wantSources []SourceConfig
	}{
		{
			name:       "add one source in empty config",
			wantSource: validSourceConfig("gh"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRepository{tt.cfg, tt.repoErr}
			service := NewService(repo)

			targetErr := service.Add(tt.wantSource)

			source, _ := service.Get(tt.wantSource.Name)

			if tt.wantErr != "" {
				assert.ErrorContains(t, targetErr, tt.wantErr)
			}

			assert.NoError(t, targetErr)
			assert.Equal(t, tt.wantSource, source)
		})
	}
}

func TestConfig_Service_Remove(t *testing.T) {
	tests := []struct {
		name        string
		cfg         Config
		nameRemove  string
		repoErr     error
		wantErr     string
		wantSource  SourceConfig
		wantSources []SourceConfig
	}{
		{
			name:        "remove exists source",
			cfg:         validConfigWithNamedSources([]string{"gh"}),
			nameRemove:  "gh",
			wantSources: []SourceConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRepository{tt.cfg, tt.repoErr}
			service := NewService(repo)

			targetErr := service.Remove(tt.nameRemove)

			sources, _ := service.List()

			if tt.wantErr != "" {
				assert.ErrorContains(t, targetErr, tt.wantErr)
			}

			assert.NoError(t, targetErr)
			assert.Equal(t, tt.wantSources, sources)
		})
	}
}
