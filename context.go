package ginlike

type Context struct {
	writemem responseWriter
	Writer   ResponseWriter
}

func (c *Context) Render(code int, r Render) {
	r.Render(c.Writer)
}

func (c *Context) reset() {
	c.Writer = &c.writemem
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Render(code, String{Format: format, Data: values})
}
