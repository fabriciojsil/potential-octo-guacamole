package service

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestRestoreCounterService(t *testing.T) {
	t.Run("Save counter state", func(t *testing.T) {
		var writer bytes.Buffer
		ts := []time.Time{
			time.Date(2017, 06, 1, 1, 1, 06, 66666, time.UTC),
		}
		expected, _ := json.Marshal(ts)

		restore := RestoreCounterService{}
		err := restore.SaveState(&writer, ts)

		if err != nil {
			t.Errorf("Mustn't return an error %d", err)
		}

		if !reflect.DeepEqual(writer.Bytes(), expected) {
			t.Errorf("The Result is diferent, expected %v Actual %v", writer, expected)
		}
	})

	t.Run("retore counter state", func(t *testing.T) {

		var reader bytes.Buffer
		ts := []time.Time{
			time.Date(2017, 06, 1, 1, 1, 06, 66666, time.UTC),
		}
		expected, _ := json.Marshal(ts)
		reader.Write(expected)

		restore := RestoreCounterService{}
		result, err := restore.RetrieveState(&reader, int64(reader.Len()))

		if err != nil {
			t.Errorf("Mustn't return an error %d", err)
		}

		if !reflect.DeepEqual(result, ts) {
			t.Errorf("The Result is diferent, expected %v Actual %v", ts, result)
		}
	})
}
