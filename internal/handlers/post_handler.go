package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveirabalsa/go-simple-crud/internal/entities"
	"github.com/oliveirabalsa/go-simple-crud/internal/services"
)

type PostHandler struct {
	postService services.PostService
	// params that will mount the handler, nothing for now
}

func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {

	post := h.postService.GetPost()
	jsonData, err := json.Marshal(post)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post = entities.Post{}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	data, err := h.postService.CreatePost(post)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data.ToJson())

	// get
}
