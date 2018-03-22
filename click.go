package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main()  {

	listener ,err := net.Listen("tcp","localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {

		con ,err := listener.Accept()

		if err != nil {
			log.Println(err)
		}

		go handleCon(con)

	}


}

func handleCon(con net.Conn) {

	defer con.Close()

	for {
		_, err := io.WriteString(con,time.Now().Format("15:04:05\n"))

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}