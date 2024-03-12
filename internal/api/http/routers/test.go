package routers

import (
	"github.com/alireza-fa/ghofle/internal/api/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func Test(router fiber.Router) {
	testHandler := handlers.NewTestHandler()

	router.Get("/", testHandler.Get)
}
