package interfaces

import (
	"myapp/internal/entities"
)

type UserRepository interface {
	GetAll() ([]entities.User, error)
}
