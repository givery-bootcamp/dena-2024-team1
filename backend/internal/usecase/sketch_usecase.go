package usecase

import (
	"context"
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

func (u *SketchUsecase) CreateSketch(ctx context.Context, file *multipart.File) error {
	return u.sketchRepository.CreateSketch(ctx, file)
}

func (u *SketchUsecase) GetSketches(ctx context.Context) ([]entity.Sketch, error) {
	return u.sketchRepository.GetAll(ctx)
}
