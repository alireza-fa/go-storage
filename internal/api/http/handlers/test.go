package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (handler *TestHandler) Get(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
