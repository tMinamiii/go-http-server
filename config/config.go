package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env  string `env:"ENV" envDefault:"development"`
	Port int    `env:"PORT" envDefault:"3333"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
