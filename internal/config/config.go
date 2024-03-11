package config

import "github.com/alireza-fa/ghofle/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
