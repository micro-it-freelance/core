package core_db

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	config "github.com/micro-it-freelance/core/config"
)

func NewDB() *sqlx.DB {
	source := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", 
	config.DB.Name, config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port)

	db, err := sqlx.Connect("pgx", source)
	if err != nil {
		panic(err)
	}
	
	return db
}
