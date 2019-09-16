package ginlike

import "net/http"

type Render interface {
	Render(http.ResponseWriter) error
}

var (
	_ Render = String{}
)
