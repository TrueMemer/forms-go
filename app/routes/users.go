package routes

import (
	"forms/app/controllers"
	"forms/app/middleware/users"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UsersRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", controllers.IndexUsers)
	r.Post("/", controllers.StoreUser)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(users.UserCtx)
		r.Get("/", controllers.GetUser)
		r.Patch("/", controllers.UpdateUser)
		r.Delete("/", controllers.DeleteUser)
	})

	return r
}