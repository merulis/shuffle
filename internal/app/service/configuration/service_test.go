package sconfiguration

import (
	"testing"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/pkg/testh"
	"github.com/stretchr/testify/assert"
)

type fakeRepository struct {
	cfg entity.Configuration
	err error
}

func (r *fakeRepository) Load() (entity.Configuration, error) {
	return r.cfg, r.err
}

func (r *fakeRepository) Save(cfg entity.Configuration) error {
	r.cfg = cfg
	return r.err
}

func TestConfig_Service_List(t *testing.T) {
	tests := []struct {
		name        string
		cfg         entity.Configuration
		repoErr     error
		wantErr     string
		wantSources []entity.Source
	}{
		{
			name:        "return sources",
			cfg:         testh.ValidConfig(),
			wantSources: testh.ValidConfig().Sources,
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
		cfg        entity.Configuration
		repoErr    error
		wantErr    string
		wantName   string
		wantSource entity.Source
	}{
		{
			name: "found source",
			cfg: testh.ValidConfigWithNamedSources(
				[]string{"gh", "nexus", "gl"},
			),
			wantName:   "nexus",
			wantSource: testh.ValidSourceConfig("nexus"),
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
		cfg         entity.Configuration
		repoErr     error
		wantErr     string
		wantSource  entity.Source
		wantSources []entity.Source
	}{
		{
			name:       "add one source in empty config",
			wantSource: testh.ValidSourceConfig("gh"),
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
		cfg         entity.Configuration
		nameRemove  string
		repoErr     error
		wantErr     string
		wantSource  entity.Source
		wantSources []entity.Source
	}{
		{
			name:        "remove exists source",
			cfg:         testh.ValidConfigWithNamedSources([]string{"gh"}),
			nameRemove:  "gh",
			wantSources: []entity.Source{},
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
