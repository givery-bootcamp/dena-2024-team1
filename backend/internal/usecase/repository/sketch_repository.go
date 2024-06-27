package repository

import (
	"mime/multipart"
	"myapp/internal/entity"

	"github.com/gin-contrib/sessions"
)

type SketchRepository interface {
	CreateSketch(file *multipart.File, session sessions.Session) error
	GetAll() ([]entity.Sketch, error)
}
