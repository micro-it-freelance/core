package dbconn

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	"github.com/micro-it-freelance/core/config"
)

func NewDBConn() *sqlx.DB {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		config.DB.Name, config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port))
	if err != nil {
		panic(err)
	}
	return db
}
