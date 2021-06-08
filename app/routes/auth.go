package routes

import (
	"forms/app/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AuthRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", controllers.Login)

	return r
}