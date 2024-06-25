package repository

import (
	"myapp/internal/entity"

	"github.com/oapi-codegen/runtime/types"
)

type SketchRepository interface {
	CreateSketch(filenanme string, file *types.File) (entity.Sketch, error)
	GetAll() ([]entity.Sketch, error)
}
