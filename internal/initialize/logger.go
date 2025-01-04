package initialize

import (
	"SunnyContest/global"
	"SunnyContest/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
