package usecase

import (
	"context"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
)

type PostUsecase struct {
	postRepository repository.PostRepository
}

func NewPostUsecase(pr repository.PostRepository) PostUsecase {
	return PostUsecase{
		postRepository: pr,
	}
}

func (u *PostUsecase) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return u.postRepository.GetAll(ctx)
}

func (u *PostUsecase) GetPost(ctx context.Context, id int) (*entity.Post, error) {
	return u.postRepository.Get(ctx, id)
}

func (u *PostUsecase) CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error) {
	return u.postRepository.CreatePost(ctx, &post)
}

func (u *PostUsecase) UpdatePost(ctx context.Context, id int, title string, body string) (*entity.Post, error) {
	return u.postRepository.UpdatePost(ctx, id, title, body)
}

func (u *PostUsecase) DeletePost(ctx context.Context, id int) error {
	return u.postRepository.DeletePost(ctx, id)
}
