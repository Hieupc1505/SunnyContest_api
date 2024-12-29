package initialize

import (
	"SunnyContest/global"
	"SunnyContest/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	//middleware -> logging -> cross -> limiter global
	//r.Use()

	//routers
	userRouter := router.RotuerGroupApp.User

	//Main group
	MainGroup := r.Group("/v1/api/")

	{
		userRouter.InitUserRouter(MainGroup)
	}

	return r
}
