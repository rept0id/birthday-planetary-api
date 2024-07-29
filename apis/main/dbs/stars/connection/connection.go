package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var user = os.Getenv("DB_STARS_POSTGRES_USER")
var password = os.Getenv("DB_STARS_POSTGRES_PASSWORD")
var dbname = os.Getenv("DB_STARS_POSTGRES_DB")
var host = os.Getenv("DB_STARS_POSTGRES_HOST")
var port = os.Getenv("DB_STARS_POSTGRES_PORT")

var connStr = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

func Connection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return conn, err
}
