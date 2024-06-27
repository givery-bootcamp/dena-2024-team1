package usecase

import (
	"mime/multipart"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"

	"github.com/gin-contrib/sessions"
)

type SketchUsecase struct {
	sketchRepository repository.SketchRepository
	userRepository   repository.UserRepository
}

func NewSketchUsecase(pr repository.SketchRepository, ur repository.UserRepository) SketchUsecase {
	return SketchUsecase{
		sketchRepository: pr,
		userRepository:   ur,
	}
}

func (u *SketchUsecase) CreateSketch(file *multipart.File, session sessions.Session) error {

	user, err := u.userRepository.GetSessionUser(session)
	if err != nil {
		return err
	}

	return u.sketchRepository.CreateSketch(file, user.ID)
}

func (u *SketchUsecase) GetSketches() ([]entity.Sketch, error) {
	return u.sketchRepository.GetAll()
}
