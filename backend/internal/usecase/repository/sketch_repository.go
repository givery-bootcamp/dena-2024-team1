package repository

import (
	"mime/multipart"
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(file *multipart.File) error
	GetAll() ([]entity.Sketch, error)
}
