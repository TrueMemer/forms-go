package main

import (
	"forms/app"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	check(err)

	println("Starting server...")

	server := app.New()

	http.HandleFunc("/", server.Router.ServeHTTP)

	err = http.ListenAndServe(":9000", nil)
	check(err)
}

func check(e error) {
	if (e != nil) {
		log.Println(e)
		os.Exit(1)
	}
}
