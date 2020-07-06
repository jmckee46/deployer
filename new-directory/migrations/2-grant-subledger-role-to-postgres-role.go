package migrations

import (
	"github.com/jmckee46/myAppTest/images-to-deploy/postgres"

	"github.com/jmckee46/deployer/flaw"
)

func GrantSubledgerRoleToPostgresRole(pgClient *postgres.Client) flaw.Flaw {
	_, err := pgClient.Exec(
		"GRANT subledger TO postgres",
	)

	if err != nil {
		return flaw.From(err).Wrap("could not grant subledger role to postgres role")
	}

	return nil
}
