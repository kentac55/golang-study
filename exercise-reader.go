package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	for {
		b = append(b, 65)
	}
	return cap(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
