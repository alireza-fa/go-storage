package main

import (
	"github.com/alireza-fa/ghofle/cmd"
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

	root.AddCommand(
		cmd.NewServer().Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command:\n%s", err)
	}
}
