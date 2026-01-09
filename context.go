package wtype

import (
	"context"
	"time"
)

type Context[T any] struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	index     int
	handler   []func(*Context[T])
	aborted   bool
	data      map[string]any
	C         T
}

func (c *Context[T]) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context[T]) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context[T]) Err() error {
	return c.ctx.Err()
}

func (c *Context[T]) Value(key any) any {
	k, ok := Assert[string](key)
	if !ok {
		return nil
	}
	v, ok := c.Get(k)
	if !ok {
		return nil
	}
	return v
}

func (c *Context[T]) clone() IContext {
	ctx, cancel := context.WithCancel(context.Background())
	cp := Context[T]{
		ctx:       ctx,
		ctxCancel: cancel,
		index:     0,
		handler:   make([]func(*Context[T]), len(c.handler)),
		aborted:   false,
		data:      make(map[string]any, len(c.data)),
		C:         c.C,
	}
	for k, v := range c.data {
		cp.data[k] = v
	}
	copy(cp.handler, c.handler)
	return &cp
}

func (c *Context[T]) Next() {
	defer c.ctxCancel()
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
	cp := c.clone()
	return cp
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
