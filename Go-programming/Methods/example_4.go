package Methods

type Buffer struct {
	Buf []byte
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Buf = append(b.Buf, p...)
	return len(p), nil
}
