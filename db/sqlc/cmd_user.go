package db

import "time"

var (
	StudentRole  = 2
	TeacherRole  = 4
	StatusActive = 0
	StatusLocked = 2
)

func NewCreateUserParams(username string, password string) CreateUserParams {
	return CreateUserParams{
		Username:     username,
		Password:     password,
		Role:         int32(StudentRole),
		Status:       int32(StatusActive),
		TokenExpried: time.Now(),
	}
}
