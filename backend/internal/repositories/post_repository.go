package repositories

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) GetAll() ([]entities.Post, error) {
	var posts []Post
	result := r.Conn.Find(posts)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", posts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return convertPostRepositoryModelToEntity(posts), nil
}

func convertPostRepositoryModelToEntity(ps []Post) []entities.Post {
	var posts []entities.Post

	for _, p := range ps {
		posts = append(posts, entities.Post{
			ID:        p.ID,
			UserID:    p.UserID,
			Title:     p.Title,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			DeletedAt: p.DeletedAt,
		})
	}
	return posts
}
