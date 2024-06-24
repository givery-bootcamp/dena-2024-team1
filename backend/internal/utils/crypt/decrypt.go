package crypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func DecryptPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
