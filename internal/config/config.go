package config

import (
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
)

type Config struct {
	Port     string
	Logger   *logger.Config
	Postgres *rdbms.Config
}
