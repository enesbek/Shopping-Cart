package db

import (
	"backend/constants"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	dbHost := os.Getenv("DBHOST")
	connectionString := fmt.Sprintf("host=%s port=5432 user=postgres dbname=postgres password=1234 sslmode=disable", dbHost)
	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	db.MustExec(constants.CREATE_PRODUCT_TABLE_QUERY)

	return db
}
