package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b < 65 || (b > 90 && b < 97) || b > 122 {
		return b
	}
	if (b > 77 && b < 91) || b > 109 {
		return byte(b - 13)
	}
	return byte(b + 13)
}

func (rot13Reader *rot13Reader) Read(p []byte) (int, error) {
	n, err := (*rot13Reader).r.Read(p)
	for i, v := range p {
		p[i] = rot13(v)
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
