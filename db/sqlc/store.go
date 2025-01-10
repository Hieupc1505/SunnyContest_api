package db

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	SYSTEM_ERROR = fmt.Errorf("system_error")
)

type Store interface {
	Querier
}

type SQLStore struct {
	db DBTX
	*Queries
}

func NewStore(conn *pgxpool.Pool) Store {
	return &SQLStore{
		Queries: New(conn),
		db:      conn,
	}
}
