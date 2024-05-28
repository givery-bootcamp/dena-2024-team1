package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostUsecase struct {
	postRepository interfaces.PostRepository
}

func NewPostUsecase(pr interfaces.PostRepository) *PostUsecase {
	return &PostUsecase{
		postRepository: pr,
	}
}

func (u *PostUsecase) GetPosts() ([]entities.Post, error) {
	return u.postRepository.GetAll()
}

func (u *PostUsecase) GetPost() (*entities.Post, error) {
	return u.postRepository.Get()
}
