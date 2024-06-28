package usecase

import (
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

func (u *PostUsecase) GetPosts() ([]entity.Post, error) {
	return u.postRepository.GetAll()
}

func (u *PostUsecase) GetPost(id int) (*entity.Post, error) {
	return u.postRepository.Get(id)
}

func (u *PostUsecase) CreatePost(post entity.Post, session sessions.Session) (*entity.Post, error) {
	user, err := u.userRepository.GetSessionUser(session)
	if err != nil {
		return nil, err
	}

	post.UserID = user.ID
	post.UserName = user.Name

	return u.postRepository.CreatePost(&post)
}

func (u *PostUsecase) UpdatePost(id int, title string, body string, session sessions.Session) (*entity.Post, error) {
	canModify, err := u.canModifyPostBySessionUser(id, session)
	if err != nil {
		return nil, err
	}
	if !canModify {
		return nil, errors.New("cannot modify post")
	}

	return u.postRepository.UpdatePost(id, title, body)
}

func (u *PostUsecase) DeletePost(id int, session sessions.Session) error {
	canModify, err := u.canModifyPostBySessionUser(id, session)
	if err != nil {
		return err
	}
	if !canModify {
		return errors.New("cannot modify post")
	}

	return u.postRepository.DeletePost(id)
}

func (u *PostUsecase) canModifyPostBySessionUser(postID int, session sessions.Session) (bool, error) {
	post, err := u.postRepository.Get(postID)
	if err != nil {
		return false, err
	}

	user, err := u.userRepository.GetSessionUser(session)
	if err != nil {
		return false, err
	}

	result := post.UserID == user.ID

	return result, nil
}
