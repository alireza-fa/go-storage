package cmd

import (
	"fmt"
	"github.com/alireza-fa/ghofle/internal/api/http"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
	"os"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (cmd *Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run Ghofle server",
		Run:   run,
	}
}

func (cmd *Server) run(trap chan os.Signal) {
	server := http.New()
	go server.Serve(8080)

	filed := fmt.Sprintf("signal trap %s", (<-trap).String())
	log.Info("existing by receiving unix signal", filed)
}
