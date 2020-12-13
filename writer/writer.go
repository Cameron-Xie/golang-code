package writer

import (
	"bufio"
	"io"
	"sync"
)

type Writer interface {
	Flush() error
	io.StringWriter
}

type writer struct {
	sync.RWMutex
	w *bufio.Writer
}

func (w *writer) Flush() error {
	return w.w.Flush()
}

func (w *writer) WriteString(s string) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	return w.w.Write([]byte(s))
}

func New(w io.Writer) Writer {
	return &writer{
		w: bufio.NewWriter(w),
	}
}
