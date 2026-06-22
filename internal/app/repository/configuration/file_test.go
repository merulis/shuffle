package rconfiguration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/pkg/testh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_FileRepository_Load(t *testing.T) {
	emptySTR := ""
	invalidJSON := "{"

	validJSONBytes, err := json.Marshal(testh.ValidConfig())
	assert.NoError(t, err)

	validJSON := string(validJSONBytes)

	tests := []struct {
		name        string
		fileContent *string
		want        entity.Configuration
		wantErr     string
	}{
		{
			name:        "file does not exist",
			fileContent: nil,
			want:        entity.Configuration{},
		},
		{
			name:        "empty file",
			fileContent: &emptySTR,
			want:        entity.Configuration{},
		},
		{
			name:        "invalid json",
			fileContent: &invalidJSON,
			want:        entity.Configuration{},
			wantErr:     "decode config",
		},
		{
			name:        "valid json",
			fileContent: &validJSON,
			want:        testh.ValidConfig(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")

			if tt.fileContent != nil {
				err := os.WriteFile(path, []byte(*tt.fileContent), 0o644)
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
		cfg  entity.Configuration
	}{
		{
			name: "saves empty config",
			cfg:  entity.Configuration{},
		},
		{
			name: "saves config with data",
			cfg:  testh.ValidConfig(),
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
