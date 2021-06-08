package dtos

import (
	"forms/app/models"
	"net/http"
)

type LoginDto struct {
	Username string
	Password string
}

type LoginResponse struct {
	*models.User
	Token string
}

func (u *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (d *LoginDto) Bind(r *http.Request) error {
	// TODO: Add validation
	return nil
}