package models

import (
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type FormType string

const (
	FORM_QUIZ FormType = "quiz"
	FORM_TEST FormType = "test"
)

type Form struct {
	gorm.Model
	Title string
	Type FormType
	Class string
	Groups []Group `gorm:"many2many:form_groups;"`
}

type FormResponse struct {
	*Form
}

func (u *FormResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewFormResponse(form *Form) *FormResponse {
	resp := &FormResponse{Form: form}

	return resp
}

func NewFormListResponse(forms []*Form) []render.Renderer {
	var list []render.Renderer
	for _, form := range forms {
		list = append(list, NewFormResponse(form))
	}
	return list
}