package tests

import (
	"net/http"
	"oneQrCode/bootstrap"
	"testing"
)

// mockNewServer.
func mockNewServer() *http.Server {
	//bootstrap.Setup()
	return bootstrap.NewServe()
}

func TestAccount(t *testing.T) {
	router := mockNewServer()
	// todo:// 还需要做的事情： 1.API地址优化  2.加入jwt认证
	items := []APITestCase{
		{
			"test_getCaptcha",
			"GET",
			"/api/v1/account/getCaptcha",
			"",
			nil,
			http.StatusOK,
			`{"request_id":"<<PRESENCE>>","code":50000,"status_code":401,"message":"token 验证失败：未识别的 payload"}`,
		},
	}

	for _, item := range items {
		Endpoint(t, router, item)
	}
}
