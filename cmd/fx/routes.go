package fx

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//type RouteGroup struct {
//	Prefix string
//	Routes []Route
//}

//type Route struct {
//	Method  string
//	Path    string
//	Handler http.HandlerFunc
//}

func AddRoutes() *chi.Mux {
	r := chi.NewRouter()

	// default middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// TODO: add validator

	// TODO: add route here

	// TODO: add fallbacks

	return r
}
