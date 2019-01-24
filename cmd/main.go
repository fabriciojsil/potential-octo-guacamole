package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fabriciojsil/potential-octo-guacamole/cmd/app"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/gateway/counter"
)

func main() {
	mux := http.NewServeMux()
	tmpFile := os.TempDir() + "/state.json"
	ticker := time.NewTicker(time.Second)
	c := counter.NewCounterRequest(time.Minute, ticker)

	application := &app.App{
		Counter:         c,
		RestorePathFile: tmpFile,
	}

	application.Start(mux)

	server := http.Server{Addr: ":3000", Handler: mux}
	log.Printf("Server is running")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
