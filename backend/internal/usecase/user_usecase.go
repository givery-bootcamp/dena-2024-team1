package usecase

import (
	"myapp/internal/usecase/repository"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: ur,
	}
}
