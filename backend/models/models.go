package models

import (
	"backend/constants"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source models/models.go -destination mocks/model_mock.go -package=mocks
type ProductModelImpl interface{
	DbGetProducts() ([]Product, error)
	DbAddProductToBasket(productId int) error
	DbGetBasketProducts() ([]Product, error)
	DbIncrementProductQuantity(productId int) error
	DbDecrementProductQuantity(productId int) error
	DbDeleteProductFromBasket(productId int) error
}

type (
	Product struct{
		Id   	 int    `json:"id"`
		Name 	 string `json:"name"`
		Price 	 int 	`json:"price"`
		Image 	 string `json:"image"`
		Quantity int	`json:"quantity"`
		Basket   bool   `json:"basket"`
	}

	ProductModel struct{
		db *sqlx.DB
	}
)

func NewProductModel(db *sqlx.DB) *ProductModel {
	return &ProductModel{
		db: db,
	}
}

func (p *ProductModel) DbGetProducts() ([]Product, error) {
	rows, err := p.db.Query(constants.GET_ALL_PRODUCTS_QUERY)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []Product
	var product Product
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Image, &product.Quantity, &product.Basket)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return products, nil
}

func (p *ProductModel) DbAddProductToBasket(productId int) error{
	_, err := p.db.Exec(constants.ADD_PRODUCT_TO_BASKET, productId)

	if err != nil {
		panic(err)
	}

	return nil
}

func (p *ProductModel) DbGetBasketProducts() ([]Product, error) {
	rows, err := p.db.Query(constants.GET_ALL_BASKET)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []Product
	var product Product
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Image, &product.Quantity, &product.Basket)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return products, nil
}

func (p *ProductModel) DbIncrementProductQuantity(productId int) error {
	_, err := p.db.Exec(constants.INCREMENT_QUANTITY_OF_PRODUCT, productId)

	if err != nil {
		panic(err)
	}

	return nil
}

func (p *ProductModel) DbDeleteProductFromBasket(productId int) error {
	_, err := p.db.Exec(constants.DELETE_PRODUCT_FROM_BASKET, productId)

	if err != nil {
		panic(err)
	}

	return nil
}

func (p *ProductModel) DbDecrementProductQuantity(productId int) error {
	var quantity int
	var err error
	p.db.QueryRow(constants.GET_QUANTITY_OF_PRODUCT, productId).Scan(&quantity)

	if quantity > 1 {
		_, err = p.db.Exec(constants.DECREMENT_QUANTITY_OF_PRODUCT, productId)

		if err != nil {
			panic(err)
		}

	} else {
		p.DbDeleteProductFromBasket(productId)
	}

	return nil
}



