package presenter

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONPresenter(t *testing.T) {
	t.Run("successful present ", func(t *testing.T) {
		expected := `{"A":1}`
		w := httptest.NewRecorder()
		p := JSONPresenter{
			Writer: w,
		}

		data := struct{ A int }{A: 1}

		p.Present(data)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("The StatusCode is diferent, expected %v Actual %v", resp.StatusCode, http.StatusOK)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		if string(body) != expected {
			t.Errorf("The result is diferent, expected %v Actual %v", string(body), expected)
		}
	})
	t.Run("Unsuccessful present ", func(t *testing.T) {
		expected := `Fail to convert count`
		w := httptest.NewRecorder()
		p := JSONPresenter{
			Writer: w,
		}

		data := make(chan int)

		p.Present(data)

		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("The StatusCode is diferent, expected %v Actual %v", resp.StatusCode, http.StatusInternalServerError)
		}

		body, _ := ioutil.ReadAll(resp.Body)

		if string(body) != expected {
			t.Errorf("The result is diferent, expected %v Actual %v", string(body), expected)
		}
	})

}
