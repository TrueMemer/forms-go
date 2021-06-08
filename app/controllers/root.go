package controllers

import (
	"fmt"
	"net/http"
)

func RootIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
