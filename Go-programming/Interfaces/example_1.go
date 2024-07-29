package Interfaces

type ByteCounter struct {
	Count int
}

// Write implements the io.Writer interface for ByteCounter.
func (c *ByteCounter) Write(p []byte) (int, error) {
	c.Count += len(p)
	return len(p), nil
}
