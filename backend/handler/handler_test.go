package unit

import (
	"Shopping-Cart/backend/handler"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var productJSON = `{"id": 1, "name"":"iphone", "price": 1000, "image": "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg"}`

func TestGetAllProducts(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/allproducts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler.Handler{}
	sampleProduct := handler.Product1
	if assert.NoError(t, h.GetAllProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		actual := recorderToJSON(rec.Body)
		assert.Equal(t, sampleProduct, actual)
	}

}

func recorderToJSON(recorder *bytes.Buffer) handler.Product {
	body, err := ioutil.ReadAll(recorder)
	if err != nil {
		fmt.Println(err)
	}
	actual := handler.Product1
	err = json.Unmarshal(body, &actual)

	if err != nil {
		fmt.Println(err)
	}
	return actual
}