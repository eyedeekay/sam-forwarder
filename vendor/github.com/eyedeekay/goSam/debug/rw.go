package debug

import (
	"encoding/hex"
	"io"
	"log"
)

/*
Copy of testing/iotest Read- and WriteLogger, but using %q instead of %x for printing
*/

type writeLogger struct {
	prefix string
	w      io.Writer
}

func (l *writeLogger) Write(p []byte) (n int, err error) {
	n, err = l.w.Write(p)
	if err != nil {
		log.Printf("%s %q: %v", l.prefix, string(p[0:n]), err)
	} else {
		log.Printf("%s %q", l.prefix, string(p[0:n]))
	}
	return
}

// NewWriteLogger returns a writer that behaves like w except
// that it logs (using log.Printf) each write to standard error,
// printing the prefix and the hexadecimal data written.
func NewWriteLogger(prefix string, w io.Writer) io.Writer {
	return &writeLogger{prefix, w}
}

type readLogger struct {
	prefix string
	r      io.Reader
}

func (l *readLogger) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p)
	if err != nil {
		log.Printf("%s %q: %v", l.prefix, string(p[0:n]), err)
	} else {
		log.Printf("%s %q", l.prefix, string(p[0:n]))
	}
	return
}

// NewReadLogger returns a reader that behaves like r except
// that it logs (using log.Print) each read to standard error,
// printing the prefix and the hexadecimal data written.
func NewReadLogger(prefix string, r io.Reader) io.Reader {
	return &readLogger{prefix, r}
}

type readHexLogger struct {
	prefix string
	r      io.Reader
}

func (l *readHexLogger) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p)
	if err != nil {
		log.Printf("%s (%d bytes) Error: %v", l.prefix, n, err)
	} else {
		log.Printf("%s (%d bytes)", l.prefix, n)
	}
	log.Print("\n" + hex.Dump(p[:n]))
	return
}

// NewReadHexLogger returns a reader that behaves like r except
// that it logs to stderr using ecoding/hex.
func NewReadHexLogger(prefix string, r io.Reader) io.Reader {
	return &readHexLogger{prefix, r}
}
