package infrastructure

import (
	"fmt"

	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {

	const (
		HOST     = "127.0.0.1"
		PORT     = 5432
		USER     = "nichola"
		PASSWORD = "123"
		DBNAME   = "nusaresto"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	db, err := sqlx.Open("postgres", connStr)

	if db.Ping() == err {
		helper.PanicIfError(err)
	}

	return db, err
}
