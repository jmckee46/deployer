package main

import (
	"os"

	awsclient "github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/myAppTest/images-to-deploy/postgres"
)

type state struct {
	AWSClient    *awsclient.Client `json:"-"`
	PGClient     *postgres.Client  `json:"-"`
	CreatedAt    string            `json:"created-at"`
	FlawError    flaw.Flaw         `json:"flaw-error"`
	File         os.FileInfo       `json:"-"`
	FileID       string            `json:"file-id"`
	FileName     string            `json:"file-name"`
	FilepathName string            `json:"filepath-name"`
	Files        []os.FileInfo     `json:"-"`
	Message      string            `json:"message"`
}

func newState() *state {

	state := &state{
		PGClient: postgres.Connect(
			"postgres",
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("DE_MIGRATIONS_PGHOST"),
		),
	}

	return state
}
