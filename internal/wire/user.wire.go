//go:build wireinject

package wire

import (
	"SunnyContest/internal/controller"
	"SunnyContest/internal/services"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		services.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
