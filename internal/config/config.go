package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
)

type WebConfig struct {
	HostPort string `env:"SERVER_ADDRESS"`
	BaseUrl  string `env:"BASE_URL"`
}

type Config struct {
	WebConfig WebConfig
}

func InitFlagConfig() *Config {
	cfg := &Config{}

	env.Parse(&cfg.WebConfig)
	if cfg.WebConfig.HostPort == "" {
		pflag.StringVar(&cfg.WebConfig.HostPort, "a", "localhost:8080", "server host")
	}
	if cfg.WebConfig.BaseUrl == "" {
		pflag.StringVar(&cfg.WebConfig.BaseUrl, "b", "http://localhost:8080", "base url")
		pflag.Parse()
	}
	return cfg
}
