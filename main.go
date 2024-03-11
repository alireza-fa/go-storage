package main

import (
	"github.com/alireza-fa/ghofle/cmd"
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const description = "Ghofle Application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.Default()

	root.AddCommand(
		cmd.NewServer().Command(cfg, trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command:\n%s", err)
	}
}
