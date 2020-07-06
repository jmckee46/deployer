package migrations

import (
	"github.com/jmckee46/myAppTest/images-to-deploy/postgres"

	"github.com/jmckee46/deployer/flaw"
)

func CreateDatabase(pgClient *postgres.Client) flaw.Flaw {
	_, err := pgClient.Exec(
		"CREATE DATABASE subledger WITH OWNER subledger",
	)

	if err != nil {
		return flaw.From(err).Wrap("could not create database")
	}

	return nil
}
