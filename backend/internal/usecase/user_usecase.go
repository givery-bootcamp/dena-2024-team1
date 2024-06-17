package usecase

import (
	"errors"
	"myapp/internal/usecase/repository"

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

	// パスワードを暗号化
	password, err := encryptPassword(password)

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
