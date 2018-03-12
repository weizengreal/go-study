package main

import (
	"os"
	"fmt"
)

type Writer interface {
	Write(b []byte) (n int, err error)
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	w = os.Stdout

	fmt.Fprintln(w,"hello world 1 \n \n")

}

