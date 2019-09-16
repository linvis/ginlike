package main

import (
	"net/http"

	"github.com/ginlike"
)

func main() {
	engine := ginlike.Default()

	engine.GET("/", home)

	engine.Run("localhost:5000")
}

func home(c *ginlike.Context) {
	c.String(http.StatusOK, "hello world")
}
