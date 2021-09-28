package tests

import (
	"net/http"
	"oneQrCode/bootstrap"
	"testing"
)

// mockNewServer.
func mockNewServer() *http.Server {
	return bootstrap.NewServe()
}

func TestAccount(t *testing.T) {
	router := mockNewServer()
	items := []APITestCase{
		{
			"test_getCaptcha",
			"GET",
			"/api/v1/account/getCaptcha",
			"",
			nil,
			http.StatusOK,
			`{"code":200,"msg":"success","data":"<<PRESENCE>>"}`,
		},
	}

	for _, item := range items {
		Endpoint(t, router, item)
	}
}
