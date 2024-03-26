package rdbms

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func New(cfg *Config, development string) (RDBMS, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.HostDebug, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error ping database:\n%s", err)
	}

	return &rdbms{db: db}, nil
}
