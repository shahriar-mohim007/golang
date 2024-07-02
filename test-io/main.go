package main

import (
	"fmt"
	"io"
	"log"
)

type DataProvider struct {
	data string
}

func (dp *DataProvider) Read(p []byte) (n int, err error) {
	copy(p, dp.data)
	return len(dp.data), io.EOF
}

func process(r io.Reader) error {
	buf := make([]byte, 100)
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return err
	}
	fmt.Println(string(buf[:n]))
	return nil
}

func main() {
	dp := &DataProvider{data: "Hello, Go!"}
	err := process(dp)
	if err != nil {
		log.Fatal(err)
	}
}
