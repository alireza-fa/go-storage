package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Server struct {
	App *fiber.App
}

func New() *Server {
	s := &Server{}

	s.App = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	return s
}

func (server *Server) Serve(port int) error {
	addr := fmt.Sprintf(":%d", port)
	if err := server.App.Listen(addr); err != nil {
		log.Errorf("error resolving server: %s", err)
		return err
	}

	return nil
}
