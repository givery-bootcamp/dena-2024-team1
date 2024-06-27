package usecase

import (
	"context"
	"errors"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
	"myapp/internal/utils/crypt"

	"github.com/gin-contrib/sessions"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: ur,
	}
}

func (u UserUsecase) Signup(ctx context.Context, session sessions.Session, username, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}

	// ユーザー名が既に存在するかチェック
	user, err := u.userRepository.GetUserByUsername(ctx, username)

	if err != nil {
		return err
	}

	if user.Name != "" {
		return errors.New("username is already taken")
	}

	// パスワードを暗号化
	password, err = crypt.EncryptPassword(password)

	if err != nil {
		return err
	}

	user, err = u.userRepository.CreateUser(ctx, username, password)

	if err != nil {
		return err
	}

	// セッションにユーザーを保存
	err = u.userRepository.SaveSession(ctx, session, user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) Signin(ctx context.Context, session sessions.Session, username, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}

	hashedPassword, err := u.userRepository.GetUserPassword(ctx, username)

	if err != nil {
		return err
	}

	if password == "" {
		return errors.New("user is not found")
	}

	if err := crypt.DecryptPassword(hashedPassword, password); err != nil {
		return errors.New("password is incorrect")
	}

	user, err := u.userRepository.GetUserByUsername(ctx, username)

	if err != nil {
		return err
	}

	// セッションにユーザーを保存
	err = u.userRepository.SaveSession(ctx, session, user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetSessionUser(ctx context.Context, session sessions.Session) (*entity.User, error) {
	user, err := u.userRepository.GetSessionUser(ctx, session)

	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, errors.New("user is not found")
	}

	return &user, nil
}

func (u UserUsecase) Signout(ctx context.Context, session sessions.Session) error {
	return u.userRepository.DeleteSessionUser(ctx, session)
}
