package migrations

import (
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/myAppTest/images-to-deploy/postgres"

	"github.com/jmckee46/deployer/flaw"
)

func CreateRole(pgClient *postgres.Client, pgPassword string) flaw.Flaw {
	logger.Info("pgPassword is", pgPassword)
	_, err := pgClient.Exec(
		"CREATE ROLE subledger LOGIN ENCRYPTED PASSWORD '" + pgPassword + "'",
	)

	if err != nil {
		return flaw.From(err).Wrap("could not create subledger role")
	}

	return nil
}
