package handlers

import (
	"github.com/alireza-fa/ghofle/internal/api/dto"
	"github.com/alireza-fa/ghofle/internal/api/validations"
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/alireza-fa/ghofle/internal/services"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	logger  logger.Logger
	service *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	cfg := config.Default()
	return &AuthHandler{
		logger:  logger.NewLogger(cfg.Logger),
		service: services.NewAuthService(cfg),
	}
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	request := dto.RegisterUser{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(logger.Validation, logger.BodyParser, errString, nil)
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	if err := validations.Validators["email"].Validate(request.Email); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := validations.Validators["username"].Validate(request.Username); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := validations.Validators["password"].Validate(request.Password); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := handler.service.RegisterUser(c, &request); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}

func (handler *AuthHandler) Verify(c *fiber.Ctx) error {
	request := dto.UserVerify{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(logger.Validation, logger.BodyParser, errString, nil)
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	if err := validations.Validators["email"].Validate(request.Email); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := validations.Validators["code"].Validate(request.Code); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	token, err := handler.service.Verify(&request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(token)
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Login api")
}
