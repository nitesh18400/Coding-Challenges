package main

import (
	"log"
	"net/http"

	"url-shortner/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	server := mux.NewRouter()

	server.HandleFunc("/shorten", handlers.ShortenUrl).Methods("POST", "OPTIONS")
	server.HandleFunc("/{shortURL}", handlers.RedirectToOriginal).Methods("GET", "OPTIONS")
	server.HandleFunc("/{shortURL}", handlers.DeleteUrl).Methods("DELETE", "OPTIONS")
	c := cors.AllowAll()

	handler := c.Handler(server)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
