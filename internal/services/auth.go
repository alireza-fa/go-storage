package services

import (
	"errors"
	"fmt"
	"github.com/alireza-fa/ghofle/internal/api/dto"
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/redis"
	"github.com/alireza-fa/ghofle/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthService struct {
	cfg *redis.Config
	log logger.Logger
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		cfg: cfg.Redis,
		log: logger.NewLogger(cfg.Logger),
	}
}

type userRegisterCacheValue struct {
	Username string
	Email    string
	Password string
	Code     string
}

func (service *AuthService) RegisterUser(c *fiber.Ctx, userRegister *dto.RegisterUser) error {
	extra := map[logger.ExtraKey]interface{}{
		logger.Username: userRegister.Username,
		logger.Email:    userRegister.Email,
	}

	r, err := redis.New[*userRegisterCacheValue](service.cfg)
	if err != nil {
		errString := fmt.Sprintf("Error while get a new redis connection, error: %s", err)
		service.log.Fatal(logger.Redis, logger.Startup, errString, extra)
		return errors.New(errString)
	}
	defer r.Close()

	key := userRegister.Email + "auth"

	cacheInfo, err := redis.Get[*userRegisterCacheValue](r.Client, key)
	if err != nil {
		errString := fmt.Sprintf("Error while get cache with key: %s", key)
		service.log.Error(logger.Redis, logger.RedisGet, errString, extra)
		return errors.New(errString)
	}
	if cacheInfo != nil {
		errString := fmt.Sprintf("You received a code less than two minutes ago, email: %s", userRegister)
		service.log.Error(logger.Auth, logger.Register, errString, extra)
		return errors.New(errString)
	}

	if err = service.checkAllowIpAddressToReceiveMail(c, r); err != nil {
		service.log.Error(logger.Auth, logger.Register, err.Error(), extra)
		return err
	}

	value := userRegisterCacheValue{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: userRegister.Password,
		Code:     utils.GenerateOtpCode(),
	}

	if err = redis.Set(r.Client, key, &value, time.Duration(120)*time.Second); err != nil {
		errString := fmt.Sprintf("Error while set cache with key: %s", key)
		service.log.Error(logger.Redis, logger.RedisSet, errString, extra)
		return errors.New(errString)
	}

	return nil
}

func (service *AuthService) checkAllowIpAddressToReceiveMail(c *fiber.Ctx, r *redis.Redis) error {
	ipAddress := utils.GetClientIp(c)
	key := ipAddress + "otp_limit"
	extra := map[logger.ExtraKey]interface{}{
		logger.IpAddress: ipAddress,
	}

	count, err := redis.Get[int](r.Client, key)
	if err != nil {
		errString := fmt.Sprintf("Error while get cache with key: %s", key)
		service.log.Error(logger.Redis, logger.RedisGet, errString, extra)
		return errors.New(errString)
	}

	if count == 0 {
		if err = redis.Set(r.Client, key, 1, time.Duration(60*60*24)*time.Second); err != nil {
			errString := fmt.Sprintf("Error while set cache with key: %s", key)
			service.log.Error(logger.Redis, logger.RedisSet, errString, extra)
			return errors.New(errString)
		}
	}

	if count > 10 {
		errString := fmt.Sprintf("You cannot receive mail until 24 hours")
		service.log.Error(logger.Auth, logger.Register, errString, extra)
		return errors.New(errString)
	}

	redis.Incr(r.Client, key)

	return nil
}
