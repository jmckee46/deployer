package postgres

import (
	"database/sql"
	"os"
	"time"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func open() *sql.DB {
	for {
		pg, err := sql.Open("postgres", "")

		if err != nil {
			logger.Warn("postgres-open", flaw.From(err).Wrap("cannot open pg: "+os.Getenv("PGHOST")))
			time.Sleep(time.Second)
			continue
		}

		return pg
	}
}
