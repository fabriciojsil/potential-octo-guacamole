package handler

import (
	"net/http"

	"github.com/fabriciojsil/potential-octo-guacamole/internal/gateway/counter"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/presenter"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/service"
)

type Handler struct {
	Counter counter.Counter
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := presenter.JSONPresenter{Writer: w}
	incrementService := service.IncrementService{
		Counter:   h.Counter,
		Presenter: p,
	}

	incrementService.Run()
}
