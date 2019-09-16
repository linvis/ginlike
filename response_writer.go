package ginlike

import (
	"io"
	"net/http"
)

type ResponseWriter interface {
	http.ResponseWriter

	Reset(http.ResponseWriter)

	Size() int

	Status() int

	WriteHeaderNow(int)

	WriteString(string)
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

var _ ResponseWriter = &responseWriter{}

func (w *responseWriter) Reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
}

func (w *responseWriter) WriteHeaderNow(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) WriteString(s string) {
	io.WriteString(w.ResponseWriter, s)
}

func (w *responseWriter) Size() int {
	return w.size
}

func (w *responseWriter) Status() int {
	return w.status
}
