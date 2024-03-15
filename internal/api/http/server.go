package http

import (
	"encoding/json"
	"fmt"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Server struct {
	app    *fiber.App
	logger logger.Logger
}

func New(log logger.Logger) *Server {
	s := &Server{logger: log}

	s.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	s.AddGroups()

	return s
}

func (server *Server) Serve(port string) error {
	addr := fmt.Sprintf(":%s", port)

	server.logger.Info(logger.Server, logger.Startup, "web server started", nil)
	if err := server.app.Listen(addr); err != nil {
		server.logger.Error(logger.Server, logger.Startup, fmt.Sprintf("error resolving server: %s", err), nil)
		<-time.After(time.Second)
		return err
	}

	return nil
}
