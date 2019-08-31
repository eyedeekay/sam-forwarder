package debug

import (
	"io"
	//"github.com/miolini/datacounter"
)

type RWC struct {
	io.Reader
	io.Writer
	c io.Closer
}

func WrapRWC(c io.ReadWriteCloser) io.ReadWriteCloser {
	rl := NewReadLogger("<", c)
	wl := NewWriteLogger(">", c)

	return &RWC{
		Reader: rl,
		Writer: wl,
		c:      c,
	}
}

func (c *RWC) Close() error {
	return c.c.Close()
}

/*
type Counter struct {
	io.Reader
	io.Writer
	c io.Closer

	Cr *datacounter.ReaderCounter
	Cw *datacounter.WriterCounter
}

func WrapCounter(c io.ReadWriteCloser) *Counter {
	rc := datacounter.NewReaderCounter(c)
	wc := datacounter.NewWriterCounter(c)

	return &Counter{
		Reader: rc,
		Writer: wc,
		c:      c,

		Cr: rc,
		Cw: wc,
	}
}

func (c *Counter) Close() error {
	return c.c.Close()
}
*/
