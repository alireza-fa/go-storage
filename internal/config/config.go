package config

import (
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
	"github.com/alireza-fa/ghofle/pkg/redis"
	"github.com/alireza-fa/ghofle/pkg/token"
)

type Config struct {
	Port        string
	Development string
	Logger      *logger.Config
	Postgres    *rdbms.Config
	Redis       *redis.Config
	Token       *token.Config
}
