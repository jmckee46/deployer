package router

import (
	"github.com/halorium/httprouter"
)

// Router contains the router
type Router struct {
	*httprouter.Router
}

// New creates a new router struct
func New() *Router {
	router := httprouter.New()

	router.UseRawPath = true

	router.PanicHandler = panicHandler

	router.NotFound = &notFoundHandler{}

	return &Router{
		router,
	}
}
