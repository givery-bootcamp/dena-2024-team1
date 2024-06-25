package repository

import (
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(filenanme string, file string) (entity.Sketch, error)
	GetAll() ([]entity.Sketch, error)
}
