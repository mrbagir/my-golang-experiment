package main

import (
	"fmt"
	"io"
	"os"
)

type MyReader interface {
	Read(p []byte) (n int, err error)
}

type StaticReader struct {
	Data string
	Pos  int
}

func (r *StaticReader) Read(p []byte) (int, error) {
	if r.Pos >= len(r.Data) {
		return 0, io.EOF
	}
	n := copy(p, r.Data[r.Pos:])
	fmt.Println(n)
	r.Pos += n
	return n, nil
}

func mainReader() {
	r := &StaticReader{Data: "Hello, Baqir!"}
	io.Copy(os.Stdout, r)
}
