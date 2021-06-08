package dtos

import (
	"forms/app/models"
	"net/http"
)

type UserStoreDto struct {
	*models.User `json:"required"`
}

func (d *UserStoreDto) Bind(r *http.Request) error {
	// TODO: Add validation
	return nil
}

type UserUpdateDto struct {
	*models.User
}

func (d *UserUpdateDto) Bind(r *http.Request) error {
	// TODO: Add validation
	return nil
}