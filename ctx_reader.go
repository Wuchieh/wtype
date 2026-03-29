package wtype

import (
	"context"
	"io"
)

type ctxReader struct {
	ctx context.Context
	r   io.Reader
}

func (c *ctxReader) Read(p []byte) (int, error) {
	select {
	case <-c.ctx.Done():
		return 0, c.ctx.Err()
	default:
		return c.r.Read(p)
	}
}

// NewCtxReader create an io.Reader that supports context.Context
func NewCtxReader(ctx context.Context, reader io.Reader) io.Reader {
	return &ctxReader{ctx: ctx, r: reader}
}
