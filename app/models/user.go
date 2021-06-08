package models

import (
	"forms/utils"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"required"`
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	MiddleName string `gorm:"not null"`
	Password string `gorm:"not null" json:"-"`
	Role string `gorm:"not null;default:users"`
}

func (u *User) ComparePassword(password string) bool {
	match, err := utils.ComparePasswordAndHash(password, u.Password)
	if err != nil {
		println(err)
		return false
	}

	return match
}

type UserResponse struct {
	*User
}

func (u *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserResponse(user *User) *UserResponse {
	resp := &UserResponse{User: user}

	return resp
}

func NewUserListResponse(users []*User) []render.Renderer {
	var list []render.Renderer
	for _, user := range users {
		list = append(list, NewUserResponse(user))
	}
	return list
}

