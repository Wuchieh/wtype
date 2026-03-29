package wtype_test

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

type testCtxReader struct {
	data      []byte
	dataIndex int
}

func (t *testCtxReader) Read(p []byte) (n int, err error) {
	if t.dataIndex >= len(t.data) {
		return 0, io.EOF
	}
	time.Sleep(time.Millisecond * 100)
	n = copy(p, t.data[t.dataIndex:t.dataIndex+1])
	t.dataIndex++
	return
}

func TestNewCtxReader(t *testing.T) {
	r := &testCtxReader{
		data:      []byte("hello world"),
		dataIndex: 0,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reader := wtype.NewCtxReader(ctx, r)
	_, err := io.ReadAll(reader)
	if err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Error(err)
		}
		return
	}

	t.Error("read data error")
}
