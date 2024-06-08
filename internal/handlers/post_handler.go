package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/oliveirabalsa/go-simple-crud/internal/entities"
	"github.com/oliveirabalsa/go-simple-crud/internal/services"
)

type PostHandler struct {
	postService services.PostService
}

func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {

	post := h.postService.GetPost()
	jsonData, err := json.Marshal(post)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Error parsing Id", http.StatusBadRequest)
		return
	}
	post, err := h.postService.GetByID(parsedId)
	if err != nil {
		http.Error(w, "Post not found", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(post.ToJson())

}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post = entities.Post{}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	data, err := h.postService.CreatePost(post)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data.ToJson())
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var post = entities.Post{}
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	parsedId, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, "Error parsing Id", http.StatusBadRequest)
		return
	}

	_, err = h.postService.GetByID(parsedId)

	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	post.Id = parsedId

	updatedPost, err := h.postService.UpdatePost(post)
	if err != nil {
		http.Error(w, "Error updating post", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(updatedPost.ToJson())
}
