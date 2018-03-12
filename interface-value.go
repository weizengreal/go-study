package main

import (
	"io"
	"fmt"
	"bytes"
)

const debug = true


func _F(out io.Writer) {
	fmt.Println("do something in function _F")

	if out != nil {
		out.Write([] byte("done \n"))
	}

}


func main() {

	var buf io.Writer

	if debug {
		buf = new(bytes.Buffer)
	}

	_F(buf)

	if debug {

	}

}

