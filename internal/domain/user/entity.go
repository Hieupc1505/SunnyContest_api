package user

import (
	"fmt"
	"strings"
)

var (
	ErrInValidUsername   = fmt.Errorf("validate_login_error.username")
	ErrUsernameTaken     = fmt.Errorf("username is already taken") //with error code 15
	ErrInvalidPassword   = fmt.Errorf("validate_login_error.username")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrPasswordIncorrect = fmt.Errorf("incorrect password")
	ErrInvalidToken      = fmt.Errorf("invalid token")
	ErrTokenExpired      = fmt.Errorf("token expired")
)

const (
	MaxLengthUsername = 16
	MinLengthUsername = 6
	MinLengthPassword = MinLengthUsername
	MaxLengthPassword = MaxLengthUsername
)

type Username string

func (us Username) String() string {
	return string(us)
}

func NewUsername(name string) (Username, error) {
	name = strings.Trim(name, " ")
	if len(name) > MaxLengthUsername || len(name) < MinLengthUsername {
		return "", ErrInValidUsername
	}

	return Username(name), nil
}

type Password string

func (ps Password) String() string {
	return string(ps)
}

func NewPassword(pass string) (Password, error) {
	if len(pass) < MinLengthPassword {
		return "", ErrInvalidPassword
	}
	return Password(pass), nil
}

type User struct {
	Username string
	Password string
}

func NewUser(username Username, password Password) User {
	return User{username.String(), password.String()}
}
