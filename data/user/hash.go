package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", errors.New("could not hash the password")
	}

	return string(passwordHash), nil
}
