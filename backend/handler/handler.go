package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type product struct {
	id		 int    `json:"id"`
	name 	 string `json:"name"`
	price 	 int 	`json:"price"`
	image 	 string `json:"image"`
	quantity int	`json:"quantity"`
}

type handler struct {

}

var prodcuts = map[int]*product{}



func GetAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, prodcuts)
}

func (h *handler)PostBasketProductIncrement(c echo.Context) error {

	return nil
}