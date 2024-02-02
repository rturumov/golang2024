package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tsis1/internal/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/players", handler.GetPlayers).Methods("GET")
	r.HandleFunc("/players/{name}", handler.GetPlayer).Methods("GET")
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	log.Println("Server is starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
