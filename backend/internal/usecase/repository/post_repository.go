package repository

import (
	"myapp/internal/entity"
)

type PostRepository interface {
	GetAll() ([]entity.Post, error)
	Get(int) (*entity.Post, error)
	CreatePost(*entity.Post) (entity.Post, error)
	UpdatePost(id int, title string, body string) (*entity.Post, error)
	Delete(int) error
}
