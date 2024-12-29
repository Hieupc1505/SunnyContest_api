package initialize

import (
	db "SunnyContest/db/sqlc"
	"SunnyContest/global"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitDb() {
	connPool, err := pgxpool.New(context.Background(), global.Config.PgDb.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL database: ", err)
	}
	global.PgDb = db.NewStore(connPool)
}
