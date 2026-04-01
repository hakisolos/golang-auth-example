package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hashpassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("Password missing")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ComparePassword(password string, hashed string) bool {
	if password == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
