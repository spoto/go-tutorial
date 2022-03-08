package countingwriter

import (
	"io"
)

type wrappedWriter struct {
	counter *int64
	parent  io.Writer
}

func (c wrappedWriter) Write(p []byte) (n int, err error) {
	n, err = c.parent.Write(p)
	*c.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var counter int64
	ww := wrappedWriter{&counter, w}
	return ww, ww.counter
}
