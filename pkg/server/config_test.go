package server

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)


func TestConfig(t *testing.T) {
	tests := []struct {
		name string
		env map[string]string
		expect *Config
		wantErr error
	}{
		{
			name: "defaults",
			env: nil,
			expect: &Config{
				GinAddress: ":8080",
				RQLiteURL: "http://localhost:4001",
			},
		},
		{
			name: "defaults",
			env: map[string]string{
				"GIN_ADDRESS": ":8081",
			},
			expect: &Config{
				GinAddress: ":8081",
				RQLiteURL: "http://localhost:4001",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.env {
				oldValue := os.Getenv(key)
				os.Setenv(key, value)
				defer os.Setenv(key, oldValue)
			}

			config, err := LoadConfig()

			if test.wantErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, test.expect, config)
		})
	}
}

