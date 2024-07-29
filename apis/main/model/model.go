package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBs struct {
	DBStarsConn *sql.DB
}

type Star struct {
	ID   int
	Name string
	LY   float64
	LYPM float64
}
