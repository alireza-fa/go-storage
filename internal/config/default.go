package config

import "github.com/alireza-fa/ghofle/pkg/logger"

func Default() *Config {
	return &Config{
		Port: 8080,
		Logger: &logger.Config{
			Logger:      "zap",
			Development: true,
			Encoding:    "console",
			Level:       logger.DebugLevel,
			FilePath:    "./logs/",
			Seq: struct {
				ApiKey  string
				BaseUrl string
				Port    string
			}{ApiKey: "aefaeaveavae", BaseUrl: "localhost", Port: "5341"},
		},
	}
}
