package services

import (
	db "SunnyContest/db/sqlc"
	"SunnyContest/global"
	"SunnyContest/internal/domain/user"
	"SunnyContest/internal/dto"
	"SunnyContest/pkg/errsx"
	"SunnyContest/response"
	"context"
	"fmt"
	"net/http"
)

type IUserService interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) (*db.SfUser, error)
	Login() error
}

type userService struct{}

func NewUserService() IUserService {
	return &userService{}
}

func (s *userService) Register(ctx context.Context, req *dto.RegisterUserRequest) (*db.SfUser, error) {
	var input struct {
		username user.Username
		password user.Password
	}

	var errs errsx.Map
	var err error

	input.username, err = user.NewUsername(req.Username)
	if err != nil {
		errs.Set("Username", err)
	}

	input.password, err = user.NewPassword(req.Password)
	if err != nil {
		errs.Set("Password", err)
	}

	if errs != nil {
		return nil, rspx.NewError(fmt.Errorf("%w", errs), rspx.InvalidData, http.StatusBadRequest)
	}

	ur := user.NewUser(input.username, input.password)

	if _, err = global.PgDb.GetUserByEmail(ctx, ur.Username); err == nil {
		return nil, rspx.NewError(nil, rspx.UserAlreadyExists, http.StatusBadRequest) //Nhận về errors code 15 -> username taken
	}

	createUserParams := db.NewCreateUserParams(ur.Username, ur.Password)
	data, err := global.PgDb.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return nil, rspx.NewError(db.SYSTEM_ERROR, rspx.System, http.StatusInternalServerError)
	}

	return &data, nil
}

func (s *userService) Login() error {
	return nil
}
