package routes

import (
	"forms/app/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ProfileRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", controllers.GetProfile)

	return r
}