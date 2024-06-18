package repository

import (
	"myapp/internal/entity"
)

type PostRepository interface {
	GetAll() ([]entity.Post, error)
	Get(int) (*entity.Post, error)
	CreatePost(*entity.Post) (entity.Post, error)
	Delete(int) error
}
