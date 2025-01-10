package controller

import (
	db "SunnyContest/db/sqlc"
	"SunnyContest/internal/dto"
	"SunnyContest/internal/services"
	rp "SunnyContest/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var req *dto.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rp.ErrorResponse(ctx, db.SYSTEM_ERROR)
		return
	}

	user, err := uc.userService.Register(ctx, req)
	if err != nil {
		rp.ErrorResponse(ctx, err)
		return
	}

	// Register
	rp.SuccessResponse(ctx, user)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req dto.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rp.ErrorResponse(ctx, db.SYSTEM_ERROR)
		return
	}
	// Login
	//ctx.JSON(200, gin.H{"login": "successfully logged in"})
	rp.SuccessResponse(ctx, gin.H{"login": "successfully logged in"})
}
