package router

import (
	"net/http"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func panicHandler(w http.ResponseWriter, r *http.Request, recovered interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	_, err := w.Write(nil)

	if err != nil {
		panic(flaw.From(err).Wrap("cannot panicHandler"))
	}

	logger.Critical("router-panic-handler", recovered)
}
