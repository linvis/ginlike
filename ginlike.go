package ginlike

import (
	"net/http"
	"sync"
)

type RouterGroup interface {
	addRouter()
}

type Engine struct {
	routerMap map[string]HandlerFunc
	pool      sync.Pool
}

type HandlerFunc func(*Context)

func Default() *Engine {
	engine := &Engine{}

	engine.routerMap = make(map[string]HandlerFunc)

	engine.pool.New = func() interface{} {
		return &Context{}
	}

	return engine
}

func (engine *Engine) GET(path string, handlers HandlerFunc) {
	engine.routerMap[path] = handlers
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := engine.routerMap[r.URL.Path]; err != nil {
		// var ctx Context
		ctx := engine.pool.Get().(*Context)
		ctx.writemem.Reset(w)
		ctx.reset()
		engine.routerMap[r.URL.Path](ctx)
		engine.pool.Put(ctx)
	}
}

func (engine *Engine) Run(path string) {
	http.ListenAndServe(path, engine)
}
