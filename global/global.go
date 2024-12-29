package global

import (
	db "SunnyContest/db/sqlc"
	"SunnyContest/pkg/logger"
	"SunnyContest/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	PgDb   db.Store
)
