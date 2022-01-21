package handler

import (
	"backend/constants"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	ProductModel models.ProductModelImpl
}

func NewHandler(ProductModel models.ProductModelImpl) *handler {
	return &handler{
		ProductModel: ProductModel}
}

func (h *handler) GetAllProducts(c echo.Context) error {
	products, err := h.ProductModel.DbGetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constants.PRODUCT_NOT_FOUND)
	}

	return c.JSON(http.StatusOK, products)
}

func (h *handler) PostAddToBasket(c echo.Context) error {
	responseId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	err = h.ProductModel.DbAddProductToBasket(responseId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constants.PRODUCT_NOT_FOUND)
	}

	return c.JSON(http.StatusOK, true)
}

func (h *handler) GetAllBasketProducts(c echo.Context) error {
	products, err := h.ProductModel.DbGetBasketProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constants.PRODUCT_NOT_FOUND)
	}

	return c.JSON(http.StatusOK, products)
}

func (h *handler) PostBasketIncrement(c echo.Context) error {
	responseId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	err = h.ProductModel.DbIncrementProductQuantity(responseId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constants.PRODUCT_NOT_FOUND)
	}

	return c.JSON(http.StatusOK, true)
}

func (h *handler) PostBasketDecrement(c echo.Context) error {
	responseId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	err = h.ProductModel.DbDecrementProductQuantity(responseId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constants.PRODUCT_NOT_FOUND)
	}

	return c.JSON(http.StatusOK, true)
}

func (h *handler) DeleteBasketProduct(c echo.Context) error {
	responseId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	err = h.ProductModel.DbDeleteProductFromBasket(responseId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constants.PRODUCT_NOT_FOUND)
	}

	return c.JSON(http.StatusOK, true)
}