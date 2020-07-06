package migrations

import (
	"database/sql"

	"github.com/jmckee46/deployer/flaw"
)

func RestrictPrivileges(transaction *sql.Tx) flaw.Flaw {
	_, err := transaction.Exec(
		"REVOKE ALL ON SCHEMA public FROM public",
	)

	if err != nil {
		transaction.Rollback()
		return flaw.From(err).Wrap("could not revoke all on public schema")
	}

	_, err = transaction.Exec(
		"REVOKE ALL ON SCHEMA subledger FROM subledger",
	)

	if err != nil {
		transaction.Rollback()
		return flaw.From(err).Wrap("could not revoke all on subledger schema")
	}

	return nil
}
