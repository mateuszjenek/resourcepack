package models

import (
	"fmt"
	"time"

	"github.com/deerling/resourcepack/internal/config"
	"github.com/dgrijalva/jwt-go"
)

// User represent logged user
type User struct {
	Username string
	Email    string
	PassHash string
	Role     UserPrivileges
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
		Subject:   "reservation.app",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.SecretKey)
	if err != nil {
		return ss, fmt.Errorf("error while signing token: %w", err)
	}
	return ss, nil
}
