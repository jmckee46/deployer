package server

import (
	"net/http"

	envvars "github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/router"
)

// ListenAndServe starts an HTTP(S) server
func ListenAndServe(router *router.Router) {
	router.GET("/health-check", getStatus)

	go func() {
		logger.Debug("server-listen-and-serve", nil)

		err := http.ListenAndServe(":80", router)

		if err != nil {
			logger.Panic(
				"server-listen-and-serve",
				flaw.From(err).Wrap("cannot ListenAndServe"),
			)
		}
	}()

	if envvars.Mocked {
		select {}
	}

	go func() {
		logger.Debug("server-listen-and-serve-tls", router)

		err := http.ListenAndServeTLS(":443", "/tls/files/certificate-chain.pem", "/tls/files/private-key.pem", router)

		if err != nil {
			logger.Panic(
				"server-listen-and-serve-tls",
				flaw.From(err).Wrap("cannot ListenAndServe"),
			)
		}
	}()

	select {}
}
