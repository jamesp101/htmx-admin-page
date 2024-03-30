package main

import (
	"log"
	"main/handlers"
	"net/http"
	"os"

	m "main/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load .env file. Using default values.")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	r.Mount("/", handlers.IndexHandler())
	r.Mount("/auth", handlers.AuthHandler())

	r.Group(func(gr chi.Router) {
		gr.Use(m.SessionMiddleware)
		gr.Mount("/dashboard", handlers.DashboardHandler())
		gr.Mount("/accounts", handlers.AccountsHandler())
	})

	// Port Setup
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is starting on port %s \n", port)
	err = http.ListenAndServe("0.0.0.0:"+port, r)
	if err != nil {
		log.Fatalf("Server error!:\n %s", err.Error())
		panic(err)
	}
}
