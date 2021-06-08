package middleware

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Database(database *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
			ctx := context.WithValue(r.Context(), "DB", database.WithContext(timeoutContext))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}