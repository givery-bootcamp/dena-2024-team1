package interfaces

import (
	"myapp/internal/entities"
)

type PostRepository interface {
	GetAll() ([]entities.Post, error)
	Get() (*entities.Post, error)
}
