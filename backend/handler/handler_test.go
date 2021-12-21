package handler

import (
	"Shopping-Cart/backend/mocks"
	"Shopping-Cart/backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProductsSuccesfully(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	 products := []models.Product{
		models.Product{
			Id:1,
			Name: "iphone",
			Price: 1000,
			Image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg",
			Quantity: 1,
			Basket: false,
		},
	}

	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)
	m.
		EXPECT().DbGetProducts().
		Return(products, nil)

	h := NewHandler(m)

	if assert.NoError(t, h.GetAllProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		actual := recorderToJSON(rec.Body)
		assert.Equal(t, products, actual)
	}
}

func TestGetAllProductsEmpty(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	var products []models.Product

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)
	m.
		EXPECT().
		DbGetProducts().
		Return(products, nil)

	h := NewHandler(m)

	if assert.NoError(t, h.GetAllProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		actual := recorderToJSON(rec.Body)
		assert.Equal(t, products, actual)
	}
}

func TestPostAddToBasketSuccesfully(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	var product = models.Product{
		Id: 1,
		Name: "samsung",
		Price: 900,
		Image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/117735-ana_large.jpg",
		Quantity: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)
	m.
		EXPECT().
		DbAddProductToBasket(product.Id).
		Return(nil)

	h := NewHandler(m)

	if assert.NoError(t, h.PostAddToBasket(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string("true\n"), rec.Body.String())
	}
}

func TestPostBasketIncrementSuccessfully(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	var product = models.Product{
		Id: 1,
		Name: "samsung",
		Price: 900,
		Image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/117735-ana_large.jpg",
		Quantity: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)
	m.
		EXPECT().
		DbIncrementProductQuantity(product.Id).
		Return(nil)

	h := NewHandler(m)

	if assert.NoError(t, h.PostBasketIncrement(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string("true\n"), rec.Body.String())
	}
}

func TestPostBasketDecrementSuccessfully(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	var product = models.Product{
		Id: 1,
		Name: "samsung",
		Price: 900,
		Image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/117735-ana_large.jpg",
		Quantity: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)
	m.
		EXPECT().
		DbDecrementProductQuantity(product.Id).
		Return(nil)

	h := NewHandler(m)

	if assert.NoError(t, h.PostBasketDecrement(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string("true\n"), rec.Body.String())
	}
}

func TestPostDeleteBasketProductSuccessfully(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	var product = models.Product{
		Id: 1,
		Name: "samsung",
		Price: 900,
		Image: "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/117735-ana_large.jpg",
		Quantity: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductModelImpl(ctrl)
	m.
		EXPECT().
		DbDeleteProductFromBasket(product.Id).
		Return(nil)

	h := NewHandler(m)

	if assert.NoError(t, h.DeleteBasketProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string("true\n"), rec.Body.String())
	}
}

func recorderToJSON(recorder *bytes.Buffer) []models.Product {
	body, err := ioutil.ReadAll(recorder)

	if err != nil {
		fmt.Println(err)
	}

	var actual []models.Product
	err = json.Unmarshal(body, &actual)

	if err != nil {
		fmt.Println(err)
	}

	return actual
}