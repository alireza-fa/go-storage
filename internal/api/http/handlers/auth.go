package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	return c.Status(http.StatusCreated).SendString("register user api")
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("login user api")
}
