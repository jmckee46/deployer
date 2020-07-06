package main

import (
	"github.com/jmckee46/deployer/router"
	"github.com/jmckee46/deployer/server"
)

func main() {
	router := router.New()

	server.ListenAndServe(router)
}
