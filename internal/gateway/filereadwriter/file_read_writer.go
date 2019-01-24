package filereadwriter

import (
	"io"
)

//FileReadWriter responsible to read and write a file into FS
type FileReadWriter struct {
	ReadWriter io.ReadWriter
}

//Write content to a file
func (f FileReadWriter) Write(data []byte) (int, error) {
	return f.ReadWriter.Write(data)
}

//Read content from a file
func (f FileReadWriter) Read(p []byte) (int, error) {
	return f.ReadWriter.Read(p)
}
