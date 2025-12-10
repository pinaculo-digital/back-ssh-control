package server

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitConnection() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	num, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, num, user, password, dbname)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		db.Close()
		return nil, err
	}

	return db, nil

}
