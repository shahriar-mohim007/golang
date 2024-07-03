package Interfaces

import "io"

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type MyReadWriteCloser struct {
	Data string
}

func (m *MyReadWriteCloser) Read(p []byte) (n int, err error) {
	if m.Data == "" {
		return 0, io.EOF
	}
	n = copy(p, m.Data)
	m.Data = m.Data[n:]
	return n, nil
}

func (m *MyReadWriteCloser) Write(p []byte) (n int, err error) {
	m.Data += string(p)
	return len(p), nil
}

func (m *MyReadWriteCloser) Close() error {
	m.Data = ""
	return nil
}
