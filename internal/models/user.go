package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jenusek/resourcepack/internal/config"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// User represent logged user
type User struct {
	Username   string
	Email      string
	PassHash   string
	Privileges UserPrivileges
}

// UserPrivileges define user's privileges
type UserPrivileges int

const (
	UserPrivilegesStandard UserPrivileges = iota
	UserPrivilegesManager
	UserPrivilegesAdmin
)

func (u *User) GenerateToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    u.Username,
		Subject:   "resourcepack",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.SecretKey)
	if err != nil {
		return ss, fmt.Errorf("error while signing token: %w", err)
	}
	return ss, nil
}

func (u *User) GeneratePassword() (string, error) {
	pass, err := password.Generate(10, 2, 3, false, true)
	if err != nil {
		return pass, fmt.Errorf("error while generating password for user: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return pass, fmt.Errorf("error while generating hash from password: %w", err)
	}
	u.PassHash = string(hash)

	return pass, nil
}
