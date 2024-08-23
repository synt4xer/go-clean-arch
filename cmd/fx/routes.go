package fx

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/synt4xer/go-clean-arch/internal/delivery/http"
)

func AddRoutes(users *http.UsersHttp) *chi.Mux {
	r := chi.NewRouter()

	// default middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// TODO: add validator

	// TODO: add route here
	r.Route("/users", func(r chi.Router) {
		r.Post("/", users.Create)       // POST /users - create a new user
		r.Get("/", users.GetAll)        // GET /users - list all users
		r.Get("/{id}", users.GetByID)   // GET /users/{id} - get a specific user by ID
		r.Put("/{id}", users.Update)    // PUT /users/{id} - update a user by ID
		r.Delete("/{id}", users.Delete) // DELETE /users/{id} - delete a user by ID
	})

	// TODO: add fallbacks

	return r
}
