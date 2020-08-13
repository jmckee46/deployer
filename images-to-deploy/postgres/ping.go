package postgres

import (
	"database/sql"
	"os"
	"time"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func ping(pg *sql.DB) {
	for {
		err := pg.Ping()

		if err != nil {
			logger.Warn("postgres-ping", flaw.From(err).Wrap("cannot ping postgres: "+os.Getenv("PGHOST")))
			time.Sleep(time.Second)
			continue
		}

		break
	}
}
