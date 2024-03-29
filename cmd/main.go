package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oliveirabalsa/go-simple-crud/internal/handlers"
	"github.com/oliveirabalsa/go-simple-crud/internal/repositories"
	"github.com/oliveirabalsa/go-simple-crud/internal/services"
)

func main() {
	postRepository := repositories.NewPostRepository()
	postService := services.NewPostService(*postRepository)
	postHandler := handlers.NewPostHandler(*postService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", postHandler.GetPost)
	r.Post("/", postHandler.CreatePost)
	http.ListenAndServe(":8082", r)
}
