package services

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return hash, fmt.Errorf("error while encrypting password: %v", err)
	}
	return hash, nil
}

func GeneratePassword() (string, error) {
	pass, err := password.Generate(10, 2, 0, false, true)
	if err != nil {
		return pass, fmt.Errorf("error while generating password for user: %v", err)
	}
	return pass, nil
}
