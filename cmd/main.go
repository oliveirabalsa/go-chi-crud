package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/oliveirabalsa/go-simple-crud/internal/config"
	"github.com/oliveirabalsa/go-simple-crud/internal/handlers"
	"github.com/oliveirabalsa/go-simple-crud/internal/repositories"
	"github.com/oliveirabalsa/go-simple-crud/internal/services"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db := config.InitDB()

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(*postRepository)
	postHandler := handlers.NewPostHandler(*postService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", postHandler.GetPost)
	r.Get("/{id}", postHandler.GetPostById)
	r.Post("/", postHandler.CreatePost)
	r.Patch("/{id}", postHandler.UpdatePost)
	http.ListenAndServe(":8082", r)
}
