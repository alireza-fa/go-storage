package config

import "github.com/alireza-fa/ghofle/pkg/logger"

func Default() *Config {
	return &Config{
		logger: &logger.Config{
			Logger:      "seq",
			Development: true,
			Encoding:    "console",
			Level:       logger.DebugLevel,
		},
	}
}
