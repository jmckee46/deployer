package main

import (
	envvars "github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/myAppTest/images-to-deploy/postgres"
)

func main() {
	logger.Info("000-create-migrations-table", logger.EmptyMessage())

	pgClient := postgres.Connect(
		"postgres",
		envvars.Get("POSTGRES_PASSWORD"),
		envvars.Get("DE_MIGRATIONS_PGHOST"),
	)

	_, err := pgClient.Exec(
		`CREATE TABLE migrations (
		  id         varchar   NOT NULL,
		  filename   varchar   NOT NULL,
		  created_at timestamp NOT NULL,
		  PRIMARY KEY (id)
	  )`,
	)

	if err != nil {
		logger.Panic("000-create-migrations-table", flaw.From(err))
	}
}
