package http

import (
	"github.com/alireza-fa/ghofle/internal/api/http/routers"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) AddGroups() {
	v1 := server.app.Group("api/v1")

	server.v1(v1)
}

func (server *Server) v1(v1 fiber.Router) {
	// Groups
	auth := v1.Group("auth")

	// routers
	routers.Auth(auth)
}
