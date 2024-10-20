package main

import (
	"chi_pgx/internal/config"
	"chi_pgx/internal/database"
	"chi_pgx/internal/domain"
	"chi_pgx/internal/jsonlog"
	"os"
	"sync"
)

type application struct {
	config     config.Config
	logger     *jsonlog.Logger
	wg         sync.WaitGroup
	Repository domain.Repository
}

func main() {
	cfg := config.LoadConfig()
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	db, err := repository.NewPgRepository(cfg.DataBaseUrl)
	if err != nil {
		logger.PrintError(err, map[string]string{
			"context": "initializing the database",
		})
		os.Exit(1)
	}

	app := &application{
		config:     cfg,
		logger:     logger,
		Repository: db,
	}

	err = app.serve()

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "serving the application",
		})
	}

}
