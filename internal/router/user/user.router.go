package user

import (
	"SunnyContest/internal/wire"
	"github.com/gin-gonic/gin"
	"log"
)

type UserRouter struct{}

func (*UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//userHandlerNonDependency, err := wire.InitUserHandler(Router)
	userHandlerNonDependency, err := wire.InitUserRouterHandler()
	if err != nil {
		log.Fatal("cannot use userHadlerNoDependency")
	}

	UserRouterPublic := Router.Group("/user")
	{
		UserRouterPublic.POST("/register", userHandlerNonDependency.Register)
		UserRouterPublic.POST("/login", userHandlerNonDependency.Login)
	}
}
