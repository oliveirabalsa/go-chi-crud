package repositories

import (
	"github.com/google/uuid"
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

func (pr *PostRepository) GetPostById(id uuid.UUID) (entities.Post, error) {
	var post entities.Post
	pr.DB.First(&post, id)
	return post, nil
}

func (pr *PostRepository) UpdatePost(post entities.Post) entities.Post {
	pr.DB.Save(&post)
	return post
}
