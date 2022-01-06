package tests

import (
	"Shopping-Cart/backend/handler"
	"Shopping-Cart/backend/mocks"
	"Shopping-Cart/backend/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func startServer(t *testing.T) {
	e :=echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	products := []models.Product{
									{
										Id:1,
										Name: "iphone",
										Price: 1000,
										Image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg",
										Quantity: 1,
										Basket: false,
									},
								}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)

	m.EXPECT().DbGetProducts().Return(products, nil)
	m.EXPECT().DbIncrementProductQuantity(products[0].Id)
	m.EXPECT().DbDecrementProductQuantity(products[0].Id)
	m.EXPECT().DbGetBasketProducts().Return(products, nil)
	m.EXPECT().DbAddProductToBasket(products[0].Id)
	m.EXPECT().DbDeleteProductFromBasket(products[0].Id)

	h := handler.NewHandler(m)

	e.GET("/quotes/allproducts", h.GetAllProducts)
	e.POST("/quotes/basket/products/increment/:id", h.PostBasketIncrement)
	e.POST("/quotes/basket/products/decrement/:id", h.PostBasketDecrement)
	e.GET("/quotes/basket", h.GetAllBasketProducts)
	e.POST("/quotes/basket/add/:id", h.PostAddToBasket)
	e.POST("/quotes/basket/delete/:id", h.DeleteBasketProduct)

	e.Logger.Fatal(e.Start(":3000"))
}

func TestProvider(t *testing.T) {
	pact := dsl.Pact{
		Provider: "ShoppingCartProvider",
		DisableToolValidityCheck: true,
		Host: "127.0.0.1",
	}

	go startServer(t)

	_, err:= pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:			"http://localhost:3000",
		PactURLs:  					[]string{"https://enesbek.pactflow.io/pacts/barApi/ShoppingCartProvider/fooApi/ShoppingCartConsumer/version/2.0.9"},
		BrokerToken:				"BB6cjlL4fdMPxwxdQS6uLA",
		ProviderVersion:			"2.0.0",
		PublishVerificationResults: 	true,
	})

	if err != nil {
		t.Fatal(err)
	}
}