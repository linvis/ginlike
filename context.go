package ginlike

type Context struct {
	writemem responseWriter
	Writer   ResponseWriter
}

func (c *Context) String(code int, values string) {
	c.writemem.WriteHeaderNow(code)
	c.writemem.WriteString(values)
}
