package groups

import (
	"context"
	"forms/app/errors"
	"forms/app/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

func GroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		db := ctx.Value("DB").(*gorm.DB)

		group := models.Group{}
		var groupId string

		if groupId = chi.URLParam(r, "id"); groupId == "" {
			render.Render(w, r, errors.ErrNotFound)
			return
		}

		tx := db.Begin()
		err := tx.First(&group, groupId).Error
		tx.Commit()

		if (err != nil) {
			render.Render(w, r, errors.ErrNotFound)
			return
		}

		ctx = context.WithValue(r.Context(), "group", group)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
