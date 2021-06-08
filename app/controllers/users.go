package controllers

import (
	"forms/app/dtos"
	"forms/app/errors"
	"forms/app/models"
	"forms/utils"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value("user").(models.User)

	render.Render(w, r, models.NewUserResponse(&user))
}

func IndexUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	var users []*models.User

	tx := db.Begin()
	tx.Find(&users)
	tx.Commit()

	render.RenderList(w, r, models.NewUserListResponse(users))
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	data := &dtos.UserStoreDto{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	hashPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		println(err)
		hashPassword = ""
	}

	data.User.Password = hashPassword

	user := data.User

	tx := db.Begin()
	tx.Create(&user)
	tx.Commit()

	render.Render(w, r, models.NewUserResponse(user))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	user := ctx.Value("user").(models.User)

	data := &dtos.UserUpdateDto{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	hashPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		println(err)
		hashPassword = ""
	}

	tx := db.Begin()
	tx.Model(&user).Updates(&models.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		MiddleName: data.MiddleName,
		Password: hashPassword,
	})
	tx.Commit()

	render.Render(w, r, models.NewUserResponse(&user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	user := ctx.Value("user").(models.User)

	tx := db.Begin()
	tx.Delete(&user)
	tx.Commit()

	render.Render(w, r, models.NewUserResponse(&user))
}