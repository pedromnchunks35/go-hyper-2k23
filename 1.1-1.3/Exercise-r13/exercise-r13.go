package main

import (
	"io"
	"os"
)

type Reader struct {
	reader io.Reader
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
			p[i] += 13
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return
}

func main() {
	// You will receive the input of a pipe this way, by passing the pointer to a variable
	r := &Reader{reader: os.Stdin}
	buffer := make([]byte, 1024)
	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
}
