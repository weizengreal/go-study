package main

import (
	"flag"
	"time"
	"fmt"
)

var period = flag.Duration("period", 1*time.Second,"sleep period")

func main() {

	flag.Parse()

	fmt.Println("sleep %v",*period)

	time.Sleep(*period)

	fmt.Println("I am weak")

}

