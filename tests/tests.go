package tests

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// APITestCase represents the data needed to describe an API test case.
type APITestCase struct {
	Name         string
	Method       string
	URL          string
	Body         string
	Header       http.Header
	WantStatus   int
	WantResponse string
}

// Endpoint tests an HTTP endpoint using the given APITestCase spec.
func Endpoint(t *testing.T, router *http.Server, tc APITestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
		if tc.Header != nil {
			req.Header = tc.Header
		}
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}
		w := httptest.NewRecorder()
		//router.ServeHTTP(w, req)
		router.Handler.ServeHTTP(w, req)
		assert.Equal(t, tc.WantStatus, w.Code, "status mismatch")
		//assert.Equal(t, tc.WantStatus, w.Body.String())
		assert.JSONEq(t, tc.WantResponse, w.Body.String())
	})
}

func MockAuthHeader(JWT string) http.Header {
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Bearer %v", JWT))
	return header
}
