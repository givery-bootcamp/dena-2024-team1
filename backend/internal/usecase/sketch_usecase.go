package usecase

import (
	"mime/multipart"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
)

type SketchUsecase struct {
	sketchRepository repository.SketchRepository
}

func NewSketchUsecase(pr repository.SketchRepository) SketchUsecase {
	return SketchUsecase{
		sketchRepository: pr,
	}
}

func (u *SketchUsecase) CreateSketch(file *multipart.File, userid int) error {
	return u.sketchRepository.CreateSketch(file, userid)
}

func (u *SketchUsecase) GetSketches() ([]entity.Sketch, error) {
	return u.sketchRepository.GetAll()
}
