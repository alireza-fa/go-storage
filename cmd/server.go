package cmd

import (
	"fmt"
	"github.com/alireza-fa/ghofle/internal/api/http"
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (cmd *Server) Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(cfg, trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run Ghofle server",
		Run:   run,
	}
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {
	log := logger.NewLogger(cfg.Logger)

	server := http.New(log)
	go server.Serve(cfg.Port)

	filed := fmt.Sprintf("signal trap %s", (<-trap).String())
	log.Info(logger.Server, logger.Startup, "existing by receiving unix signal", map[logger.ExtraKey]interface{}{logger.Signal: filed})
}
