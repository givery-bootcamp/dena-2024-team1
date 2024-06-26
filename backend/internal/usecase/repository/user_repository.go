package repository

import (
	"myapp/internal/entity"

	"github.com/gin-contrib/sessions"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	CreateUser(username, password string) error
	GetUserByUsername(username string) (entity.User, error)
	SaveSession(session sessions.Session, user entity.User) error
	GetSessionUser(session sessions.Session) (entity.User, error)
	DeleteSessionUser(session sessions.Session) error
	GetUserPassword(username string) (string, error)
}
