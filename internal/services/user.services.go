package services

import (
	db "SunnyContest/db/sqlc"
	"SunnyContest/global"
	"SunnyContest/internal/dto"
	"SunnyContest/response"
	"context"
	"fmt"
)

type IUserService interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) (db.SfUser, error)
	Login() error
}

type userService struct{}

func NewUserService() IUserService {
	return &userService{}
}

func (s *userService) Register(ctx context.Context, req *dto.RegisterUserRequest) (db.SfUser, error) {
	fmt.Println("Username:::", req.Username)
	if _, err := global.PgDb.GetUserByEmail(ctx, req.Username); err == nil {
		return db.SfUser{}, response.WrapError(err, "UserAlreadyExists", "")
	}

	user, err := global.PgDb.CreateUser(context.Background(), db.CreateUserParams{
		Username:     "HIuec cpcy",
		Password:     "HIuec cpcy",
		Role:         1,
		Status:       1,
		Token:        "Hieupcdt",
		TokenExpried: 23253464556,
	})

	if err != nil {
		return db.SfUser{}, response.WrapError(err, "System", "SystemError")
	}

	return user, nil
}

func (s *userService) Login() error {
	return nil
}
