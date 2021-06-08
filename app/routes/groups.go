package routes

import (
	"forms/app/controllers"
	"forms/app/middleware/groups"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GroupsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", controllers.IndexGroups)
	r.Post("/", controllers.StoreGroup)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(groups.GroupCtx)
		r.Get("/", controllers.GetGroup)
		r.Patch("/", controllers.UpdateGroup)
		r.Delete("/", controllers.DeleteGroup)
	})

	return r
}