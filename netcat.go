package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	con, err := net.Dial("tcp","localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()

	copyMust(os.Stdout,con)

}

func copyMust(write io.Writer, read io.Reader) {

	_,err := io.Copy(write,read)

	if err != nil {
		log.Fatal(err)
	}

}