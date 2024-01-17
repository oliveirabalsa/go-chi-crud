package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oliveirabalsa/go-simple-crud/internal/handlers"
	"github.com/oliveirabalsa/go-simple-crud/internal/services"
)

func main() {
	postService := services.NewPostService()
	postHandler := handlers.NewPostHandler(*postService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", postHandler.GetPost)
	http.ListenAndServe(":8080", r)
}
