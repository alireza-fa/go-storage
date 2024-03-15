package cmd

import (
	"github.com/alireza-fa/ghofle/internal/config"
	"github.com/alireza-fa/ghofle/internal/db"
	"github.com/alireza-fa/ghofle/internal/db/models"
	"github.com/alireza-fa/ghofle/pkg/logger"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
	"github.com/spf13/cobra"
	"os"
)

type Migrate struct{}

func (m Migrate) Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		m.run(cfg, args)
	}

	return &cobra.Command{
		Use:       "migrate",
		Short:     "run migrations",
		Run:       run,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
	}
}

func (m Migrate) run(cfg *config.Config, args []string) {
	log := logger.NewLogger(cfg.Logger)

	if len(args) != 1 {
		log.Fatal(logger.RDBMS, logger.Migration, "Invalid arguments given", map[logger.ExtraKey]interface{}{"Args": args})
	}

	rd, err := rdbms.New(cfg.Postgres, cfg.Development)
	if err != nil {
		log.Fatal(logger.RDBMS, logger.Migration, "Error creating rdbms", map[logger.ExtraKey]interface{}{logger.Error: err.Error()})
	}

	if err := db.Migrate(models.Migrate(args[0]), rd); err != nil {
		log.Fatal(logger.RDBMS, logger.Migration, "Error while migrating", map[logger.ExtraKey]interface{}{logger.Error: err.Error(), "direction": args[0]})
	}

	log.Info(logger.RDBMS, logger.Migration, "Database has been migrated successfully", map[logger.ExtraKey]interface{}{"direction": args[0]})
}
