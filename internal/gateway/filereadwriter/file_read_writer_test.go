package filereadwriter

import (
	"bytes"
	"reflect"
	"testing"
)

func TestFileReadWriter(t *testing.T) {
	t.Run("write File", func(t *testing.T) {
		expected := "hey there :)"
		var buffer bytes.Buffer

		f := FileReadWriter{
			ReadWriter: &buffer,
		}

		f.Write([]byte(expected))

		if buffer.String() != expected {
			t.Errorf("The result is diferent, expected %v Actual %v", expected, buffer.String())
		}
	})

	t.Run("read File", func(t *testing.T) {
		expected := "hey there :)"

		var buffer bytes.Buffer
		byteExpected := []byte(expected)
		buffer.Write(byteExpected)

		fileWriter := FileReadWriter{
			ReadWriter: &buffer,
		}

		result := make([]byte, len(byteExpected))
		_, err := fileWriter.Read(result)
		if err != nil {
			t.Errorf("The result is diferent, expected %v Actual %v", expected, result)
		}

		if !reflect.DeepEqual(result, []byte(expected)) {
			t.Errorf("The result is diferent, expected %v Actual %v", expected, result)
		}
	})

}
