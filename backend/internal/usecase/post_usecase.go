package usecase

import (
	"context"
	"errors"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"

	"github.com/gin-contrib/sessions"
)

type PostUsecase struct {
	postRepository repository.PostRepository
	userRepository repository.UserRepository
}

func NewPostUsecase(pr repository.PostRepository, ur repository.UserRepository) PostUsecase {
	return PostUsecase{
		postRepository: pr,
		userRepository: ur,
	}
}

func (u *PostUsecase) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return u.postRepository.GetAll(ctx)
}

func (u *PostUsecase) GetPost(ctx context.Context, id int) (*entity.Post, error) {
	return u.postRepository.Get(ctx, id)
}

func (u *PostUsecase) CreatePost(ctx context.Context, post entity.Post, session sessions.Session) (*entity.Post, error) {
	user, err := u.userRepository.GetSessionUser(ctx, session)
	if err != nil {
		return nil, err
	}

	post.UserID = user.ID
	post.UserName = user.Name

	return u.postRepository.CreatePost(ctx, &post)
}

func (u *PostUsecase) UpdatePost(ctx context.Context, id int, title string, body string, session sessions.Session) (*entity.Post, error) {
	canModify, err := u.canModifyPostBySessionUser(ctx, id, session)
	if err != nil {
		return nil, err
	}
	if !canModify {
		return nil, errors.New("cannot modify post")
	}

	return u.postRepository.UpdatePost(ctx, id, title, body)
}

func (u *PostUsecase) DeletePost(ctx context.Context, id int, session sessions.Session) error {
	canModify, err := u.canModifyPostBySessionUser(ctx, id, session)
	if err != nil {
		return err
	}
	if !canModify {
		return errors.New("cannot modify post")
	}

	return u.postRepository.DeletePost(ctx, id)
}

func (u *PostUsecase) canModifyPostBySessionUser(ctx context.Context, postID int, session sessions.Session) (bool, error) {
	post, err := u.postRepository.Get(ctx, postID)
	if err != nil {
		return false, err
	}

	user, err := u.userRepository.GetSessionUser(ctx, session)
	if err != nil {
		return false, err
	}

	result := post.UserID == user.ID

	return result, nil
}
