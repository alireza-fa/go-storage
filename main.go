package main

import (
	"github.com/alireza-fa/ghofle/cmd"
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}
}

func main() {
	const description = "Go Storage Application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.Default()

	root.AddCommand(
		cmd.NewServer().Command(cfg, trap),
		cmd.Migrate{}.Command(cfg, trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command:\n%s", err)
	}
}
