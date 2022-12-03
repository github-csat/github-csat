package server

import (
	"github.com/codingconcepts/env"
	"github.com/pkg/errors"
)

type Config struct {
	GinAddress string `env:"GIN_ADDRESS"`
	RQLiteURL string `env:"RQLITE_URL"`
}

func LoadConfig() (*Config, error) {
	config := Config{
		GinAddress: ":8080",
		RQLiteURL: "http://localhost:4001",
	}

	if err := env.Set(&config); err != nil {
		return nil, errors.Wrap(err, "load env config")
	}
	return &config, nil
}