package server

import (
	"github.com/codingconcepts/env"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/url"
)

type Config struct {
	GinAddress string `env:"GIN_ADDRESS"`
	RQLiteURL  string `env:"RQLITE_URL"`
	StaticDir  string `env:"STATIC_DIR"`

	ProxyFrontend    string `env:"PROXY_FRONTEND"`
	ProxyFrontendURL *url.URL

	GitHubClientID     string `env:"GITHUB_CLIENT_ID"`
	GitHubClientSecret string `env:"GITHUB_CLIENT_SECRET"`

	GitHubEndpoint oauth2.Endpoint
}

func LoadConfig() (*Config, error) {
	config := Config{
		GinAddress:    ":8080",
		RQLiteURL:     "http://localhost:4001?disableClusterDiscovery=true",
		ProxyFrontend: "http://localhost:5173",

		GitHubEndpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	if err := env.Set(&config); err != nil {
		return nil, errors.Wrap(err, "load env config")
	}

	proxyURL, err := url.Parse(config.ProxyFrontend)
	if err != nil {
		return nil, errors.Wrap(err, "load env config")
	}

	config.ProxyFrontendURL = proxyURL
	return &config, nil
}
