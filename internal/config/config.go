package config

import "github.com/alireza-fa/ghofle/pkg/logger"

type Config struct {
	Port   int            `koanf:"port"`
	Logger *logger.Config `koanf:"logger"`
}
