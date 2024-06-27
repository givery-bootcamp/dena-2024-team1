package repository

import (
	"mime/multipart"
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(file *multipart.File, userID int) (entity.Sketch, error)
	GetAll() ([]entity.Sketch, error)
}
