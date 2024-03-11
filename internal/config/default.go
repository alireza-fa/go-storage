package config

import "github.com/alireza-fa/ghofle/pkg/logger"

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Logger:      "dummy",
			Development: true,
			Encoding:    "console",
			Level:       logger.DebugLevel,
			Seq: struct {
				ApiKey  string "koanf:\"api_key\""
				BaseUrl string "koanf:\"base_url\""
				Port    string "koanf:\"port\""
			}{ApiKey: "aefaeaveavae", BaseUrl: "localhost", Port: "5341"},
		},
	}
}
