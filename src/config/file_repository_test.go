package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_FileRepository_Load(t *testing.T) {
	emptySTR := ""
	invalidJSON := "{"

	validJSONBytes, err := json.Marshal(validConfig())
	assert.NoError(t, err)

	validJSON := string(validJSONBytes)

	tests := []struct {
		name        string
		fileContent *string
		want        Config
		wantErr     string
	}{
		{
			name:        "file does not exist",
			fileContent: nil,
			want:        Config{},
		},
		{
			name:        "empty file",
			fileContent: &emptySTR,
			want:        Config{},
		},
		{
			name:        "invalid json",
			fileContent: &invalidJSON,
			want:        Config{},
			wantErr:     "decode config",
		},
		{
			name:        "valid json",
			fileContent: &validJSON,
			want:        validConfig(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")

			if tt.fileContent != nil {
				err := os.WriteFile(path, []byte(*tt.fileContent), 0644)
				assert.NoError(t, err)
			}

			repo := NewFileRepository(path)
			got, err := repo.Load()

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFileRepository_Save(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
	}{
		{
			name: "saves empty config",
			cfg:  Config{},
		},
		{
			name: "saves config with data",
			cfg:  validConfig(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "nested", "config.json")
			repo := NewFileRepository(path)

			err := repo.Save(tt.cfg)
			require.NoError(t, err)

			require.FileExists(t, path)

			got, err := repo.Load()
			require.NoError(t, err)

			require.Equal(t, tt.cfg, got)
		})
	}
}

func validConfig() Config {
	return Config{
		Sources: []SourceConfig{
			{
				Name:  "gh-cli",
				Type:  SourceTypeGithub,
				Owner: "cli",
				Repo:  "cli",
				Ref:   "trunk",
			},
		},
		Default: "gh-cli",
	}
}
