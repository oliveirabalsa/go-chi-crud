package services

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/go-simple-crud/internal/entities"
	"github.com/oliveirabalsa/go-simple-crud/internal/repositories"
)

type PostService struct {
	postRepository repositories.PostRepository
}

func NewPostService(postRepository repositories.PostRepository) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}

func (ps *PostService) GetPost() []entities.Post {
	return ps.postRepository.GetPosts()
}

func (ps *PostService) CreatePost(post entities.Post) (entities.Post, error) {
	post.Id = uuid.New()
	return ps.postRepository.CreatePost(post), nil

}
