package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - computes password hash,
// returns password hash and error status
func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", errors.New("could not hash the password")
	}

	return string(passwordHash), nil
}

// IsCorrectPassword - checks whether the provided password hash
// mathes with the one stored in the database
// return true if yes, false otherwise
func IsCorrectPassword(user *User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	return err == nil
}
