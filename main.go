package main

import (
	"commitr/db"
	"commitr/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()

	db.Init()

	mux.HandleFunc("/", handlers.HandleHome)

	err := http.ListenAndServe(":3333", mux)
	log.Fatal(err)
}
