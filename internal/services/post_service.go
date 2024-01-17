package services

import "github.com/oliveirabalsa/go-simple-crud/internal/entities"

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func GetPost() entities.Post {
	post := entities.Post{
		Title:       "Lorem impsum",
		Description: "Lorem ipmsum iamet",
	}

	return post

}
