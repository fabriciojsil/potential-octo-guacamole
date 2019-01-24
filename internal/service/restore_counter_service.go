package service

import (
	"encoding/json"
	"io"
	"time"

	"github.com/fabriciojsil/potential-octo-guacamole/internal/gateway/filereadwriter"
)

type RestoreCounterService struct{}

//RetrieveState restore from a file a slice of time.Time
func (rc RestoreCounterService) RetrieveState(reader io.ReadWriter, size int64) ([]time.Time, error) {
	fullfill := make([]byte, size)
	frw := filereadwriter.FileReadWriter{
		ReadWriter: reader,
	}

	_, err := frw.Read(fullfill)
	if err != nil {
		return nil, err
	}

	t := []time.Time{}
	json.Unmarshal(fullfill, &t)
	return t, nil
}

//SaveState save a slice of time.Time into a file
func (rc RestoreCounterService) SaveState(writer io.ReadWriter, ts []time.Time) error {

	if len(ts) == 0 {
		return nil
	}
	data, err := json.Marshal(ts)

	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}
