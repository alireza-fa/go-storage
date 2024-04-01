package services

import (
	"errors"
	"fmt"
	"github.com/alireza-fa/ghofle/internal/api/dto"
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/redis"
	"github.com/alireza-fa/ghofle/pkg/token"
	"github.com/alireza-fa/ghofle/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthService struct {
	cfg   *config.Config
	log   logger.Logger
	token *token.Token
}

func NewAuthService(cfg *config.Config) *AuthService {
	log := logger.NewLogger(cfg.Logger)
	tokenService, err := token.New(cfg.Token)
	if err != nil {
		log.Fatal(logger.Token, logger.Startup, "cannot get a new instance of token", nil)
		panic(err)
	}
	return &AuthService{
		cfg:   cfg,
		log:   log,
		token: tokenService,
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

	r, err := redis.New(service.cfg.Redis)
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
		errString := fmt.Sprintf("You received a code less than two minutes ago, email: %s", userRegister.Email)
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

	logMessage := fmt.Sprintf("send a code for %s, code is: %s", value.Email, value.Code)
	service.log.Info(logger.Auth, logger.Register, logMessage, extra)

	return nil
}

func (service *AuthService) checkAllowIpAddressToReceiveMail(c *fiber.Ctx, r *redis.Redis) error {
	ipAddress := utils.GetClientIp(c)
	key := ipAddress + "otp_limit"
	keyShortLimit := ipAddress + "short_limit"
	extra := map[logger.ExtraKey]interface{}{
		logger.IpAddress: ipAddress,
	}

	cacheInfo, err := redis.Get[int](r.Client, keyShortLimit)
	if err != nil {
		errString := fmt.Sprintf("Error while get cache with key: %s", key)
		service.log.Error(logger.Redis, logger.RedisGet, errString, extra)
		return errors.New(errString)
	}
	if cacheInfo == 0 {
		if err = redis.Set[int](r.Client, keyShortLimit, 1, time.Duration(120)*time.Second); err != nil {
			errString := fmt.Sprintf("Error while set cache with key: %s", key)
			service.log.Error(logger.Redis, logger.RedisSet, errString, extra)
			return errors.New(errString)
		}
	} else if cacheInfo != 0 {
		errString := fmt.Sprintf("You received a code less than two minutes ago")
		service.log.Error(logger.Auth, logger.Register, errString, extra)
		return errors.New(errString)
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

func (service *AuthService) Verify(userVerify *dto.UserVerify) (*dto.UserToken, error) {
	key := userVerify.Email + "auth"
	extra := map[logger.ExtraKey]interface{}{
		logger.Email: userVerify.Email,
		logger.Code:  userVerify.Code,
	}

	r, err := redis.New(service.cfg.Redis)
	if err != nil {
		errString := fmt.Sprintf("Error while get a new redis connection, error: %s", err)
		service.log.Error(logger.Redis, logger.Startup, errString, extra)
		return nil, errors.New(errString)
	}

	ok, err := service.checkAllowToTryAgain(r, userVerify.Email)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid code")
	}

	userRegisterCache, err := redis.Get[*userRegisterCacheValue](r.Client, key)
	if err != nil {
		errString := fmt.Sprintf("Error while get userRegisterCacheValue, key: %s", key)
		service.log.Error(logger.Redis, logger.RedisGet, errString, extra)
		return nil, err
	}

	if userRegisterCache == nil {
		return nil, errors.New("invalid code")
	}

	if userRegisterCache.Code != userVerify.Code {
		return nil, errors.New("invalid code")
	}

	accessData := map[string]interface{}{
		"email":    userRegisterCache.Email,
		"username": userRegisterCache.Username,
	}
	refreshData := map[string]interface{}{
		"id": 1,
	}

	tokenData, err := service.token.CreateTokenString(accessData, refreshData)
	if err != nil {
		errString := "error while get token for user"
		service.log.Error(logger.Token, logger.GenerateToken, errString, nil)
		return nil, errors.New(errString)
	}

	userToken := dto.UserToken{
		RefreshToken: tokenData.RefreshToken,
		AccessToken:  tokenData.AccessToken,
	}

	return &userToken, nil
}

func (service *AuthService) checkAllowToTryAgain(r *redis.Redis, email string) (bool, error) {
	keyTry := email + "try"

	count, err := redis.Get[int](r.Client, keyTry)
	if err != nil {
		errString := fmt.Sprintf("Error while get cache, key: %s", keyTry)
		service.log.Error(logger.Redis, logger.RedisGet, errString, nil)
		return false, err
	}

	if count == 0 {
		count = 1
		if err = redis.Set[int](r.Client, keyTry, count, time.Duration(120)*time.Second); err != nil {
			errString := fmt.Sprintf("Error while set cache, key: %s, value: %v", keyTry, count)
			service.log.Error(logger.Redis, logger.RedisSet, errString, nil)
			return false, err
		}
	}

	if count > 5 {
		return false, nil
	}

	redis.Incr(r.Client, keyTry)

	return true, nil
}
