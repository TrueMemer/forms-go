package controllers

import (
	"forms/app/dtos"
	"forms/app/errors"
	"forms/app/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

func GetGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	group := ctx.Value("group").(models.Group)

	render.Render(w, r, models.NewGroupResponse(&group))
}

func IndexGroups(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	var groups []*models.Group

	tx := db.Begin()
	tx.Find(&groups)
	tx.Commit()

	render.RenderList(w, r, models.NewGroupListResponse(groups))
}

func StoreGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	data := &dtos.GroupStoreDto{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	group := data.Group
	
	tx := db.Begin()
	tx.Create(&group)
	tx.Commit()

	render.Render(w, r, models.NewGroupResponse(group))
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	user := ctx.Value("group").(models.Group)

	data := &dtos.GroupUpdateDto{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	tx := db.Begin()
	tx.Model(&user).Updates(&models.Group{
		Name: data.Name,
		Abbreviation: data.Abbreviation,
		Year: data.Year,
		Number: data.Number,
	})
	tx.Commit()

	render.Render(w, r, models.NewGroupResponse(&user))
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("DB").(*gorm.DB)

	group := ctx.Value("group").(models.Group)

	tx := db.Begin()
	tx.Delete(&group)
	tx.Commit()

	render.Render(w, r, models.NewGroupResponse(&group))
}