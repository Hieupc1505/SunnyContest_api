package initialize

import (
	"SunnyContest/global"
	"fmt"
)

func Run() {
	LoadConfig("./configs/")
	InitLogger()
	InitDb()

	r := InitRouter()
	port := fmt.Sprintf(":%d", global.Config.Server.Port)
	r.Run(port)
}
