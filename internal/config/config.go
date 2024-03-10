package config

import "github.com/alireza-fa/ghofle/pkg/logger"

type Config struct {
	logger *logger.Config `koanf:"logger"`
}
