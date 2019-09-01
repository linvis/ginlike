package ginlike_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ginlike"
)

func TestHealthCheckHandler(t *testing.T) {
	engine := ginlike.Default()

	engine.GET("/", func(ctx *ginlike.Context) {
		fmt.Fprintf(ctx.W, "hello world")
	})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	engine.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "hello world"
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
