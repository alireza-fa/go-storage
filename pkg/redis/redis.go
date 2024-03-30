package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	Client *redis.Client
	cfg    *Config
}

func New[T any](cfg *Config) (*Redis, error) {
	redisInstance := &Redis{cfg: cfg}
	if err := redisInstance.initRedis(cfg); err != nil {
		return nil, err
	}

	return redisInstance, nil
}

func (r *Redis) initRedis(cfg *Config) error {
	r.Client = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password:           cfg.Password,
		DB:                 cfg.Db,
		DialTimeout:        time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:        time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(cfg.WriteTimeout) * time.Second,
		PoolSize:           cfg.PoolSize,
		PoolTimeout:        time.Duration(cfg.PoolTimeout),
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: time.Duration(cfg.IdleCheckFrequency) * time.Millisecond,
	})

	_, err := r.Client.Ping().Result()
	if err != nil {
		errString := fmt.Sprintf("error ping redis, error: %s", err.Error())
		return errors.New(errString)
	}

	return nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}

func Set[T any](client *redis.Client, key string, value T, expire time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		errString := fmt.Sprintf("error while set key: %s and value: %v, error: %s", key, value, err)
		return errors.New(errString)
	}

	client.Set(key, v, expire)
	return nil
}

func Get[T any](client *redis.Client, key string) (T, error) {
	var value T = *new(T)

	v, err := client.Get(key).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			errString := fmt.Sprintf("error while get key: %s, error: %s", key, err.Error())
			return value, errors.New(errString)
		}
		return value, nil
	}

	err = json.Unmarshal([]byte(v), &value)
	if err != nil {
		errString := fmt.Sprintf("error while unmarshal redis result, key: %s, result: %v, error: %s", key, v, err.Error())
		return value, errors.New(errString)
	}

	return value, nil
}

func Incr(client *redis.Client, key string) {
	client.Incr(key)
}
