package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(store []byte) (int, error) {
	n, err := rot.r.Read(store)
	for i, c := range store {
		switch {
		case 'A' <= c && c <= 'Z':
			store[i] = (c-'A'+13)%26 + 'A'
		case 'a' <= c && c <= 'z':
			store[i] = (c-'a'+13)%26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
