package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabriciojsil/potential-octo-guacamole/internal/mock"
)

func TestHander(t *testing.T) {

	t.Run("Resquest ", func(t *testing.T) {
		expectedBody := `{"Count":1}`
		h := Handler{
			Counter: &mock.CounterFake{},
		}
		req := httptest.NewRequest("GET", "http://example.com/whatever", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		response := w.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("The StatusCode is diferent, expected %v Actual %v", response.StatusCode, http.StatusOK)
		}

		body, _ := ioutil.ReadAll(response.Body)

		if string(body) != expectedBody {
			t.Errorf("The Body is diferent, expected %v Actual %v", expectedBody, string(body))
		}
	})
}
