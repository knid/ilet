package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/knid/ilet/internal/database"
	handlers "github.com/knid/ilet/internal/handlers/http"
)

func main() {
	listenAddr, ok := os.LookupEnv("ILET_API_LISTEN_ADDRESS")
	if !ok {
		log.Println("ERROR(env): ILET_API_LISTEN_ADDRESS not defined. Using default: 0.0.0.0")
		listenAddr = "0.0.0.0"
	}
	apiListenPort, ok := os.LookupEnv("ILET_API_LISTEN_PORT")
	if !ok {
		log.Println("ERROR(env): ILET_API_LISTEN_PORT not defined. Using default: 8080")
		apiListenPort = "8080"
	}
	routerListenPort, ok := os.LookupEnv("ILET_ROUTER_LISTEN_PORT")
	if !ok {
		log.Println("ERROR(env): ILET_ROUTER_LISTEN_PORT not defined. Using default: 8081")
		routerListenPort = "8081"
	}

	db := database.PostgresDB{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Address:  os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT"),
		SSLMode:  os.Getenv("POSTGRES_SSL"),
	}

	if err := db.Connect(); err != nil {
		log.Fatal("ERROR(DB): Connecting Error: ", err)
	}

	// if err := db.CheckConnection(); err != nil {
	// 	log.Fatalf("ERROR(DB): %+v", err)
	// }

	httpHandler := handlers.HTTPHandler{
		DB: &db,
	}

	apiR := chi.NewRouter()
	routerR := chi.NewRouter()

	apiR.Use(middleware.Logger)
	apiR.Use(middleware.Recoverer)
	apiR.Use(middleware.Timeout(30 * time.Second))

	routerR.Use(middleware.Logger)
	routerR.Use(middleware.Recoverer)
	routerR.Use(middleware.Timeout(30 * time.Second))

	apiR.Route("/links", func(r chi.Router) {
		r.Get("/", httpHandler.GetAllLinks)       // Get all links
		r.Get("/{id}", httpHandler.GetLink)       // Get a link
		r.Put("/{id}", httpHandler.UpdateLink)    // Edit a link
		r.Delete("/{id}", httpHandler.DeleteLink) // Delete a link
		r.Post("/", httpHandler.CreateLink)       // Create a link
	})
	apiR.Route("/user", func(r chi.Router) {
		r.Get("/me", httpHandler.GetUser)             // Get user detail
		r.Put("/me", httpHandler.UpdateUser)          // Edit user
		r.Delete("/me", httpHandler.DeleteUser)       // Delete user
		r.Post("/register", httpHandler.RegisterUser) // Register a user
		r.Post("/login", httpHandler.LoginUser)       // Login user
	})

	routerR.Get("/{shortLink}", httpHandler.RouteToLongURL)

	apiAddr := listenAddr + ":" + apiListenPort
	routerAddr := listenAddr + ":" + routerListenPort

	go func() {
		log.Println("Router HTTP listener starting on " + routerAddr)
		log.Fatal(http.ListenAndServe(routerAddr, routerR))
	}()

	log.Println("API HTTP listener starting on " + apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr, apiR))

}
