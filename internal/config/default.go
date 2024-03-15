package config

import (
	"github.com/alireza-fa/ghofle/internal/constants"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
	"os"
)

func Default() *Config {
	return &Config{
		Port: os.Getenv(constants.PORT),
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
			Host:      os.Getenv(constants.DbHost),
			HostDebug: os.Getenv(constants.DbHostDebug),
			Port:      os.Getenv(constants.DbPort),
			Username:  os.Getenv(constants.DbUsername),
			Password:  os.Getenv(constants.DbPassword),
			Database:  os.Getenv(constants.DbDatabase),
		},
	}
}
