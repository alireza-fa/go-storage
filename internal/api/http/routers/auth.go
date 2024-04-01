package routers

import (
	"github.com/alireza-fa/ghofle/internal/api/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router) {
	handler := handlers.NewAuthHandler()

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
	router.Post("/verify", handler.Verify)
}
