package services

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/go-simple-crud/internal/entities"
)

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func (ps *PostService) GetPost() entities.Post {
	post := entities.Post{
		Title:       "Lorem impsum",
		Description: "Lorem ipmsum iamet",
	}
	return post
}

func (ps *PostService) CreatePost(post entities.Post) (entities.Post, error) {
	// do something
	post.Id = uuid.New()
	return post, nil
}
