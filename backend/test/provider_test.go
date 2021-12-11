package tests

import (
	"testing"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"github.com/pact-foundation/pact-go/types"
)

func startServer() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":3000"))
}

func TestProvider(t *testing.T) {
	pact := dsl.Pact{
		Provider: "ShoppingCartProvider",
		Consumer: "ShoppingCartConsumer",
	}

	go startServer()

	err, _ := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:	"http://localhost:3000",
		BrokerURL:			"https://enesbek.pactflow.io",
		BrokerToken:		"BB6cjlL4fdMPxwxdQS6uLA",
		ProviderVersion:	"1.0.0",
		PublishVerificationResults: true,
	})

	if err != nil {
		t.Fatal(err)
	}

}