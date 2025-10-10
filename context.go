package wtype

type Context[T any] struct {
	index   int
	handler []func(*Context[T])
	aborted bool
	data    map[string]any
	C       T
}

func (c *Context[T]) Next() {
	if c.IsAborted() {
		return
	}

	for c.index < len(c.handler) {
		if c.IsAborted() {
			return
		}
		handler := c.handler[c.index]
		c.index++
		handler(c)
	}
}

func (c *Context[T]) Abort() {
	c.aborted = true
}

func (c *Context[T]) IsAborted() bool {
	return c.aborted
}

func (c *Context[T]) Get(s string) (any, bool) {
	if c.data == nil {
		return nil, false
	}

	v, ok := c.data[s]
	return v, ok
}

func (c *Context[T]) Set(s string, a any) {
	if c.data == nil {
		c.data = make(map[string]any)
	}
	c.data[s] = a
}

func (c *Context[T]) DoBefore() IContext {
	cp := *c
	cp.index = 0
	cp.aborted = false
	cp.data = nil
	copy(cp.handler, c.handler)
	return &cp
}

func (c *Context[T]) Do() {
	ContextDo(c)
}

func NewContext[T any](c T) Context[T] {
	return Context[T]{
		C: c,
	}
}

func AddHandler[T any](ctx Context[T], handlers ...func(*Context[T])) Context[T] {
	ctx.handler = append(ctx.handler, handlers...)
	return ctx
}

func ContextDo(ctx IContext) {
	_c, ok := ctx.(IContextDoBefore)
	if ok {
		ctx = _c.DoBefore()
	}
	ctx.Next()
}
