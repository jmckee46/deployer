package migrations

import (
	"database/sql"

	"github.com/jmckee46/deployer/flaw"
)

func CreateSchema(transaction *sql.Tx) flaw.Flaw {
	_, err := transaction.Exec(
		"CREATE SCHEMA subledger AUTHORIZATION subledger",
	)

	if err != nil {
		return flaw.From(err).Wrap("could not create schema")
	}

	return nil
}
