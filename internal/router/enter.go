package router

import "SunnyContest/internal/router/user"

type RouterGroup struct {
	User user.UserRouterGroup
}

var RotuerGroupApp = new(RouterGroup)
