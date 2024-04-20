package repositories

import (
	"github.com/oliveirabalsa/go-simple-crud/internal/entities"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

func (pr *PostRepository) CreatePost(post entities.Post) entities.Post {
	pr.DB.Create(&post)
	return post
}

func (pr *PostRepository) GetPosts() []entities.Post {
	var posts []entities.Post

	pr.DB.Find(&posts)

	return posts
}
