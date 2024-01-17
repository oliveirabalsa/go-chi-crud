package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oliveirabalsa/go-simple-crud/internal/services"
)

type PostHandler struct {
	postService services.PostService
	// params that will mount the handler, nothing for now
}

// function to generate the handler, always type * and & with the params
//also inside ()
//example
// func NewAdvertisementHandler(advertisementService service.AdvertisementService) *AdvertisementHandler {
// 	return &AdvertisementHandler{
// 		advertisementService: advertisementService,
// 	}
// }

func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {

	post := services.GetPost()
	jsonData, err := json.Marshal(post)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
