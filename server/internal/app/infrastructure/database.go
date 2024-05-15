package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("failed to load .env file")
	}

	var (
		DB_HOST     = os.Getenv("DB_HOST")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		helper.PanicIfError(err)
	}

	if err = db.Ping(); err != nil {
		log.Println("failed to connect to database")
		helper.PanicIfError(err)
	}

	log.Println("successfully connected to the database")

	return db, err
}
