package ginlike

type Context struct {
	writemem responseWriter
	Writer   ResponseWriter
}

func (c *Context) Render(code int, r Render) {
	r.Render(c.Writer)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	// c.writemem.WriteHeaderNow(code)
	// c.writemem.WriteString(values)
	c.Render(code, String{Format: format, Data: values})
}
