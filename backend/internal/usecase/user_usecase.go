package usecase

import (
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

func (u UserUsecase) Signup(session sessions.Session, username, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}

	// ユーザー名が既に存在するかチェック
	user, err := u.userRepository.GetUserByUsername(username)

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

	user, err = u.userRepository.CreateUser(username, password)

	if err != nil {
		return err
	}

	// セッションにユーザーを保存
	err = u.userRepository.SaveSession(session, user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) Signin(session sessions.Session, username, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}

	hashedPassword, err := u.userRepository.GetUserPassword(username)

	if err != nil {
		return err
	}

	if password == "" {
		return errors.New("user is not found")
	}

	if err := crypt.DecryptPassword(hashedPassword, password); err != nil {
		return errors.New("password is incorrect")
	}

	user, err := u.userRepository.GetUserByUsername(username)

	if err != nil {
		return err
	}

	// セッションにユーザーを保存
	err = u.userRepository.SaveSession(session, user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetSessionUser(session sessions.Session) (*entity.User, error) {
	user, err := u.userRepository.GetSessionUser(session)

	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, errors.New("user is not found")
	}

	return &user, nil
}

func (u UserUsecase) Signout(session sessions.Session) error {
	return u.userRepository.DeleteSessionUser(session)
}
