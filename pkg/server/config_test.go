package server

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"net/url"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	parsedURL, _ := url.Parse("http://localhost:5173")

	tests := []struct {
		name    string
		env     map[string]string
		expect  *Config
		wantErr error
	}{
		{
			name: "defaults",
			env:  nil,
			expect: &Config{
				GinAddress:       ":8080",
				RQLiteURL:        "http://localhost:4001?disableClusterDiscovery=true",
				ProxyFrontend:    "http://localhost:5173",
				ProxyFrontendURL: parsedURL,
				SessionJWTSecret: "this-is-a-fake-jwt-secret",
				GitHubEndpoint: oauth2.Endpoint{
					AuthURL:  "https://github.com/login/oauth/authorize",
					TokenURL: "https://github.com/login/oauth/access_token",
				},
			},
		},
		{
			name: "change things",
			env: map[string]string{
				"GIN_ADDRESS": ":8081",
				"RQLITE_URL":  "http://rqlite.com:4001",
			},
			expect: &Config{
				GinAddress:       ":8081",
				RQLiteURL:        "http://rqlite.com:4001",
				ProxyFrontend:    "http://localhost:5173",
				ProxyFrontendURL: parsedURL,
				SessionJWTSecret: "this-is-a-fake-jwt-secret",
				GitHubEndpoint: oauth2.Endpoint{
					AuthURL:  "https://github.com/login/oauth/authorize",
					TokenURL: "https://github.com/login/oauth/access_token",
				},
			},
		},
		{
			name: "GitHub credentials",
			env: map[string]string{
				"GITHUB_CLIENT_ID":     "github-client-id",
				"GITHUB_CLIENT_SECRET": "github-client-secret",
			},
			expect: &Config{
				GinAddress:         ":8080",
				RQLiteURL:          "http://localhost:4001?disableClusterDiscovery=true",
				ProxyFrontend:      "http://localhost:5173",
				ProxyFrontendURL:   parsedURL,
				GitHubClientID:     "github-client-id",
				GitHubClientSecret: "github-client-secret",
				SessionJWTSecret:   "this-is-a-fake-jwt-secret",
				GitHubEndpoint: oauth2.Endpoint{
					AuthURL:  "https://github.com/login/oauth/authorize",
					TokenURL: "https://github.com/login/oauth/access_token",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.env {
				oldValue := os.Getenv(key)
				os.Setenv(key, value)
				if oldValue != "" {
					//goland:noinspection GoDeferInLoop
					defer os.Setenv(key, oldValue)
				} else {
					//goland:noinspection GoDeferInLoop
					defer os.Unsetenv(key)
				}
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
