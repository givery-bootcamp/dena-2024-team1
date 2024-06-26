package repository

import (
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(destination string) error
	GetAll() ([]entity.Sketch, error)
}
