package limitreader

import "io"

type wrappedReader struct {
	parent io.Reader
	n      int64
}

func (wr wrappedReader) Read(p []byte) (n int, err error) {
	read, err := wr.parent.Read(p[:n])
	n -= read
	return read, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return wrappedReader{r, n}
}
