package controllers

import (
	"forms/app/models"
	"github.com/go-chi/render"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(models.User)

	render.Render(w, r, models.NewUserResponse(&user))
}