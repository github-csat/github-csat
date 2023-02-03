package server

import (
	"github.com/codingconcepts/env"
	"github.com/pkg/errors"
	"net/url"
)

type Config struct {
	GinAddress string `env:"GIN_ADDRESS"`
	RQLiteURL  string `env:"RQLITE_URL"`
	StaticDir  string `env:"STATIC_DIR"`

	ProxyFrontend    string `env:"PROXY_FRONTEND"`
	ProxyFrontendURL *url.URL
}

func LoadConfig() (*Config, error) {
	config := Config{
		GinAddress:    ":8080",
		RQLiteURL:     "http://localhost:4001?disableClusterDiscovery=true",
		ProxyFrontend: "http://localhost:3000",
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
