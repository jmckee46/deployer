package router

import (
	"net/http"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

type notFoundHandler struct{}

func (h *notFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	_, err := w.Write(nil)

	if err != nil {
		panic(flaw.From(err).Wrap("cannot ServeHTTP"))
	}

	logger.Debug("router-not-found-handler", r.URL)
}
