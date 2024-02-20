package repositories

import "github.com/oliveirabalsa/go-simple-crud/internal/entities"

type PostRepository struct{}

var posts []entities.Post

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (pr *PostRepository) CreatePost(post entities.Post) entities.Post {
	posts = append(posts, post)
	return post
}

func (pr *PostRepository) GetPosts() []entities.Post {
	return posts
}
