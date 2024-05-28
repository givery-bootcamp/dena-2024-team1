package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostUsecase struct {
	userRepository interfaces.UserRepository
	postRepository interfaces.PostRepository
}

func NewPostUsecase(ur interfaces.UserRepository, pr interfaces.PostRepository) *PostUsecase {
	return &PostUsecase{
		userRepository: ur,
		postRepository: pr,
	}
}

func (u *PostUsecase) GetPosts() ([]entities.Post, error) {
	return u.postRepository.GetAll()
}
