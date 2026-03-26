package config

import "github.com/caarlos0/env/v11"

type Config struct {
	PostgresDSN    string `env:"POSTGRES_DSN"`
	ValkeyAddr     string `env:"VALKEY_ADDR"`
	ValkeyPassword string `env:"VALKEY_PASSWORD"`
	Port           string `env:"PORT"`
}

func New(opts *env.Options) (*Config, error) {
	cfg := new(Config)
	if opts != nil {
		if err := env.ParseWithOptions(cfg, *opts); err != nil {
			return nil, err
		}

		return cfg, nil
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
