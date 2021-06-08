package models

import (
	"fmt"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type Group struct {
	gorm.Model
	Number string
	Code string
	Name string
	Abbreviation string
	Year string
}

func (group *Group) ShortName() string {
	return fmt.Sprintf("%d-%s%s",
		group.Number,
		group.Abbreviation,
		group.Year)
}


type GroupResponse struct {
	*Group
}

func (u *GroupResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewGroupResponse(user *Group) *GroupResponse {
	resp := &GroupResponse{Group: user}

	return resp
}

func NewGroupListResponse(users []*Group) []render.Renderer {
	var list []render.Renderer
	for _, user := range users {
		list = append(list, NewGroupResponse(user))
	}
	return list
}