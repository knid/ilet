package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	listenAddr, ok := os.LookupEnv("ILET_API_LISTEN_ADDRESS")
	if !ok {
		log.Println("ERROR(env): ILET_API_LISTEN_ADDRESS not defined. Using default: 0.0.0.0")
		listenAddr = "0.0.0.0"
	}
	listenPort, ok := os.LookupEnv("ILET_API_LISTEN_PORT")
	if !ok {
		log.Println("ERROR(env): ILET_API_LISTEN_PORT not defined. Using default: 8080")
		listenPort = "8080"
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/{shortLink}", func(w http.ResponseWriter, r *http.Request) {})
	r.Route("/links", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})        // Get all links
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {})    // Get a link
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {})    // Edit a link
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {}) // Delete a link
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {})       // Create a link
	})
	r.Route("/user", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})          // Get user detail
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {})          // Edit user
		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {})       // Delete user
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {}) // Register a user
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})    // Login user
	})

	addr := listenAddr + ":" + listenPort

	log.Println("HTTP listener starting on " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
