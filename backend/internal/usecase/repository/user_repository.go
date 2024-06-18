package repository

import (
	"myapp/internal/entity"
	"net/http"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	CreateUser(username, password string) error
	GetUserByUsername(username string) (entity.User, error)
	SaveSession(r *http.Request, w http.ResponseWriter, user entity.User) error
}
