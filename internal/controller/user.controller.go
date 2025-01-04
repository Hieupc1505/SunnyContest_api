package controller

import (
	"SunnyContest/internal/dto"
	"SunnyContest/internal/services"
	rp "SunnyContest/response"
	"fmt"
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
		ctx.JSON(200, gin.H{"error": "Invalid request"})
		return
	}

	user, err := uc.userService.Register(ctx, req)
	if err != nil {
		rsp, status := rp.HandleErrorResponse(err)
		fmt.Println(rsp, status)
		ctx.JSON(status, rsp)
		return
	}

	// Register
	ctx.JSON(200, rp.SuccessResponseData(user))
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req dto.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(200, gin.H{"error": "Invalid request"})
		return
	}
	// Login
	ctx.JSON(200, gin.H{"login": "successfully logged in"})
}
