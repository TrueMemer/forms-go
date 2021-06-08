package controllers

import (
	errs "errors"
	"forms/app/dtos"
	"forms/app/errors"
	"forms/app/models"
	"forms/utils"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	data := &dtos.LoginDto{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	user := &models.User{}

	tx := db.Begin()
	result := tx.First(user, &models.User{
		Username: data.Username,
	}).Error

	if result != nil {
		render.Render(w, r, errors.ErrInvalidRequest(errs.New("Invalid credentials")))
		return
	}

	passwordMatch := user.ComparePassword(data.Password)
	if !passwordMatch {
		render.Render(w, r, errors.ErrInvalidRequest(errs.New("Invalid credentials")))
		return
	}

	token, err := utils.CreateJWTToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		render.Render(w, r, errors.ErrInternalServerError(errs.New("Failed to create JWT token")))
		return
	}

	response := dtos.LoginResponse{
		User: user,
		Token: token,
	}

	render.Render(w, r, &response)
}