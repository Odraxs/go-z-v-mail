package fixtures

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateEmailSearchMockServer(t *testing.T, response []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/emails/_search" {
			t.Errorf("Expected to request '/emails/_search', got: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}))
}
