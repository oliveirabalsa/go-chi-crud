package services

import (
	"fmt"

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

func (ps *PostService) GetByID(id uuid.UUID) (entities.Post, error) {
	post, err := ps.postRepository.GetPostById(id)
	if err != nil {
		return entities.Post{}, err
	}
	if post == (entities.Post{}) {
		return entities.Post{}, fmt.Errorf("post with ID %s not found", id)
	}
	fmt.Println(post)
	return post, nil
}

func (ps *PostService) UpdatePost(post entities.Post) (entities.Post, error) {
	return ps.postRepository.UpdatePost(post), nil
}
