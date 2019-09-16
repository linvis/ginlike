package ginlike

import (
	"fmt"
	"io"
	"net/http"
)

type String struct {
	Format string
	Data   []interface{}
}

func (r String) Render(w http.ResponseWriter) error {
	return WriteString(w, r.Format, r.Data)
}

func WriteString(w http.ResponseWriter, format string, data []interface{}) error {
	if len(data) > 0 {
		_, err := fmt.Fprintln(w, format, data)
		return err
	}

	_, err := io.WriteString(w, format)
	return err
}
