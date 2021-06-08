package middleware

import (
	"context"
	"forms/app/errors"
	"forms/app/models"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		db := ctx.Value("DB").(*gorm.DB)

		_, claims, _ := jwtauth.FromContext(ctx)

		userId := claims["user_id"]

		user := models.User{}

		tx := db.Begin()
		err := tx.First(&user, userId).Error
		tx.Commit()

		if err != nil {
			render.Render(w, r, errors.ErrNotFound)
			return
		}

		ctx = context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
