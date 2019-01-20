package presenter

import (
	"encoding/json"
	"net/http"
)

type JSONPresenter struct {
	Writer http.ResponseWriter
}

func (j JSONPresenter) Present(data interface{}) {
	body, err := json.Marshal(data)

	if err != nil {
		j.Writer.WriteHeader(http.StatusInternalServerError)
		j.Writer.Write([]byte("Fail to convert count"))
		return
	}
	j.Writer.WriteHeader(http.StatusOK)
	j.Writer.Write(body)
}
