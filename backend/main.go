package main

import (
	"Shopping-Cart/backend/db"
	"Shopping-Cart/backend/handler"
	"Shopping-Cart/backend/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e :=echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	d := db.Connect()
	productModel := models.NewProductModel(d)
	h := handler.NewHandler(productModel)

	e.GET("/api/allproducts", h.GetAllProducts)
	e.POST("/api/basket/add/:id", h.PostAddToBasket)
	e.GET("/api/basket", h.GetAllBasketProducts)
	e.POST("/api/basket/products/increment/:id", h.PostBasketIncrement)
	e.POST("/api/basket/products/decrement/:id", h.PostBasketDecrement)
	e.POST("/api/basket/delete/:id", h.DeleteBasketProduct)

	e.Logger.Fatal(e.Start(":3000"))
}