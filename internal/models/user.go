package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (u *User) GenerateToken(secretKey string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    u.Username,
		Subject:   "resourcepack",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ss, fmt.Errorf("error while signing token: %v", err)
	}
	return ss, nil
}
