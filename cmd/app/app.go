package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fabriciojsil/potential-octo-guacamole/internal/gateway/counter"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/handler"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/service"
)

//App Struct to initializate the application
type App struct {
	Counter         counter.Counter
	RestorePathFile string
}

//Start Starts the application
func (a *App) Start(mux *http.ServeMux) {
	restore := &service.RestoreCounterService{}
	if err := a.restoreState(restore); err != nil {
		log.Printf("It was not possible to restore the previous state %d", err)
	}

	h := handler.Handler{Counter: a.Counter}
	a.waitForExit(restore)
	mux.Handle("/", h)
}

func (a *App) openFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
}

func (a *App) waitForExit(restore *service.RestoreCounterService) {
	exit := make(chan os.Signal, 2)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go a.keepState(restore, exit)
}

func (a *App) restoreState(restore *service.RestoreCounterService) error {
	file, err := a.openFile(a.RestorePathFile)

	if err != nil {
		return err
	}
	defer file.Truncate(0)
	info, err := file.Stat()
	if err != nil {
		return err
	}

	ts, err := restore.RetrieveState(file, info.Size())
	if err != nil {
		return err
	}
	a.Counter.SetState(ts)
	return nil
}

func (a *App) keepState(restore *service.RestoreCounterService, exit <-chan os.Signal) {
	<-exit
	log.Printf("Preparing restore file")

	a.Counter.StopExpire()

	writeFile, err := a.openFile(a.RestorePathFile)
	if err != nil {
		log.Printf("It was not possible save a restore file %d", err)
		os.Exit(0)
		return
	}

	defer writeFile.Close()

	err = restore.SaveState(writeFile, a.Counter.Values())

	if err != nil {
		log.Printf("It was not possible save state %d", err)
		os.Exit(0)
		return
	}

	os.Exit(0)
}
