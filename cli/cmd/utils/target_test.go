package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCMD_Target(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantSource string
		wantPath   string
		wantErr    string
	}{
		{
			name:       "succes parse",
			input:      "gh-cli:/path",
			wantSource: "gh-cli",
			wantPath:   "path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseTarget(tt.input, ":")

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantSource, result.SourceName)
			assert.Equal(t, tt.wantPath, result.Path)
		})
	}
}
