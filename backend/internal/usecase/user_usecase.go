package usecase

import (
	"errors"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: ur,
	}
}

func (u UserUsecase) Signup(username, password string) error {
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
	password, err = encryptPassword(password)

	if err != nil {
		return err
	}

	return u.userRepository.CreateUser(username, password)
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to encrypt password")
	}

	return string(hash), nil
}

func (u UserUsecase) Signin(username, password string, r *http.Request, w http.ResponseWriter) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}

	password, err := u.userRepository.GetUserPassword(username)

	if err != nil {
		return err
	}

	if password == "" {
		return errors.New("user is not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password)); err != nil {
		return errors.New("password is incorrect")
	}

	user, err := u.userRepository.GetUserByUsername(username)

	if err != nil {
		return err
	}

	// セッションにユーザーを保存
	err = u.userRepository.SaveSession(r, w, user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetSessionUser(r *http.Request) (*entity.User, error) {
	user, err := u.userRepository.GetSessionUser(r)

	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, errors.New("user is not found")
	}

	return &user, nil
}

func (u UserUsecase) Signout(r *http.Request, w http.ResponseWriter) error {
	return u.userRepository.DeleteSessionUser(r, w)
}
