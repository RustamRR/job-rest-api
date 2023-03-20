package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var CouldNotHashPassword error = errors.New("could not hash password")

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", CouldNotHashPassword
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
