package repositories

import (
	"errors"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type Post struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`

	gorm.Model
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) GetAll() ([]entities.Post, error) {
	var posts []Post
	postResult := r.Conn.Find(&posts)
	if postResult.Error != nil {
		if errors.Is(postResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, postResult.Error
	}
	var users []User
	userResult := r.Conn.Find(&users)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, userResult.Error
	}
	return convertPostRepositoryModelToEntity(posts, users), nil
}

func convertPostRepositoryModelToEntity(ps []Post, us []User) []entities.Post {
	var posts []entities.Post

	for _, p := range ps {
		post := entities.Post{
			ID:        int(p.ID),
			Title:     p.Title,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		for _, u := range us {
			if p.UserID == int(u.ID) {
				post.UserName = u.Name
				break
			}
		}
		posts = append(posts, post)
	}
	return posts
}
