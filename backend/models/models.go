package models

import (
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen --destination=./mocks/modelimpl.go main ProductModelImpl
type ProductModelImpl interface{
	Temp()
}

type (
	Product struct{
		id		 int    `json:"id"`
		name 	 string `json:"name"`
		price 	 int 	`json:"price"`
		image 	 string `json:"image"`
		quantity int	`json:"quantity"`
	}

	ProductModel struct{
		db *sqlx.DB
	}
)

func Temp(){

}



