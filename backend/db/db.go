package db

import (
	"Shopping-Cart/backend/constants"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://postgres:1234@localhost/demo?sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.MustExec(constants.CREATE_PRODUCT_TABLE_QUERY)

	return db
}
