package controller

import (
	"SunnyContest/internal/services"
	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var req RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(200, gin.H{"error": "Invalid request"})
		return
	}

	// Register
	ctx.JSON(200, gin.H{"register": "successfully registered"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(200, gin.H{"error": "Invalid request"})
		return
	}

	// Login
	ctx.JSON(200, gin.H{"login": "successfully logged in"})
}
