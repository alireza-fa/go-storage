package config

import (
	"fmt"
	"github.com/alireza-fa/ghofle/internal/constants"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
	"github.com/alireza-fa/ghofle/pkg/redis"
	"os"
	"strconv"
	"sync"
)

var once sync.Once

var cfg *Config

func Default() *Config {
	once.Do(func() {
		cfg = &Config{
			Port:        os.Getenv(constants.PORT),
			Development: os.Getenv(constants.Development),
			Logger: &logger.Config{
				Logger:      os.Getenv(constants.LoggerName),
				Development: os.Getenv(constants.Development),
				Encoding:    os.Getenv(constants.ZapEncoding),
				Level:       os.Getenv(constants.LogLevel),
				FilePath:    os.Getenv(constants.ZapFilePath),
				Seq: struct {
					ApiKey  string
					BaseUrl string
					Port    string
				}{ApiKey: os.Getenv(constants.SeqApiKey), BaseUrl: os.Getenv(constants.SeqBaseUrl), Port: os.Getenv(constants.SeqPort)},
			},
			Postgres: &rdbms.Config{
				Host:     os.Getenv(constants.DbHost),
				Port:     os.Getenv(constants.DbPort),
				Username: os.Getenv(constants.DbUsername),
				Password: os.Getenv(constants.DbPassword),
				Database: os.Getenv(constants.DbDatabase),
			},
			Redis: &redis.Config{
				Host:               os.Getenv(constants.RedisHost),
				Port:               os.Getenv(constants.RedisPort),
				Password:           os.Getenv(constants.RedisPassword),
				Db:                 asciiToInteger(constants.RedisDb),
				DialTimeout:        asciiToInteger(constants.DialTimeout),
				ReadTimeout:        asciiToInteger(constants.ReadTimeout),
				WriteTimeout:       asciiToInteger(constants.WriteTimeout),
				PoolSize:           asciiToInteger(constants.PoolSize),
				PoolTimeout:        asciiToInteger(constants.PoolTimeout),
				IdleCheckFrequency: asciiToInteger(constants.IdleCheckFrequency),
			},
		}
		if cfg.Development == "true" {
			cfg.Postgres.Host = os.Getenv(constants.DbHostDebug)
			cfg.Redis.Host = os.Getenv(constants.RedisDebugHost)
		}
	})

	return cfg
}

func asciiToInteger(environment string) int {
	en, err := strconv.Atoi(os.Getenv(environment))
	if err != nil {
		panic(fmt.Sprintf("error while getting %s", environment))
	}
	return en
}
