package app

import (
	"fmt"
	"net/http"
)

func (app *App) IndexHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}
}
