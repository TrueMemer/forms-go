package dtos

import (
	"forms/app/models"
	"net/http"
)

type GroupStoreDto struct {
	*models.Group `json:"required"`
}

func (d *GroupStoreDto) Bind(r *http.Request) error {
	// TODO: Add validation
	return nil
}

type GroupUpdateDto struct {
	*models.Group
}

func (d *GroupUpdateDto) Bind(r *http.Request) error {
	// TODO: Add validation
	return nil
}