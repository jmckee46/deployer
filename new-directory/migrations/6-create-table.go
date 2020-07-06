package migrations

import (
	"database/sql"

	"github.com/jmckee46/deployer/flaw"
)

func CreateTable(transaction *sql.Tx) flaw.Flaw {
	_, err := transaction.Exec(
		`CREATE TABLE entities (
			prefix varchar NOT NULL,
			id varchar NOT NULL,
			etag varchar NOT NULL,
			created_at timestamp NOT NULL,
			effective_at timestamp,
			header jsonb NOT NULL,
			body jsonb NOT NULL,
			PRIMARY KEY (prefix, id, etag)
		)`,
	)

	if err != nil {
		transaction.Rollback()
		return flaw.From(err).Wrap("could not create table")
	}

	_, err = transaction.Exec(
		`CREATE INDEX created_at_index ON entities (created_at)`,
	)

	_, err = transaction.Exec(
		`CREATE INDEX effective_at_index ON entities (effective_at)`,
	)

	if err != nil {
		transaction.Rollback()
		return flaw.From(err).Wrap("could not created_at index")
	}

	return nil
}
