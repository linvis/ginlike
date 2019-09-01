package ginlike

import (
	"net/http"
)

type RouterGroup interface {
	addRouter()
}

type Engine struct {
	routerMap map[string]HandlerFunc
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

type HandlerFunc func(*Context)

func Default() *Engine {
	engine := &Engine{}

	engine.routerMap = make(map[string]HandlerFunc)

	return engine
}

func (engine *Engine) GET(path string, handlers HandlerFunc) {
	engine.routerMap[path] = handlers
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := engine.routerMap[r.URL.Path]; err != nil {
		var ctx Context
		ctx.W = w
		ctx.R = r
		engine.routerMap[r.URL.Path](&ctx)
	}
}

func (engine *Engine) Run(path string) {
	http.ListenAndServe(path, engine)
}
