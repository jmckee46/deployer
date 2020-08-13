package postgres

import "os"

// Connect the postgres connection
func Connect(
	pgUser,
	pgPassword,
	pgHost string) *Client {

	os.Setenv("PGUSER", pgUser)
	os.Setenv("PGPASSWORD", pgPassword)
	os.Setenv("PGHOST", pgHost)

	pg := open()

	ping(pg)

	return &Client{pg}
}
