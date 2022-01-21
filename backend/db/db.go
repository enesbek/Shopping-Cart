package db

import (
	"backend/constants"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=1234 sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.MustExec(constants.CREATE_PRODUCT_TABLE_QUERY)

	return db
}
