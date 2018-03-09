package main

import (
	"fmt"
	"time"
)

func main() {

	sum := 0

	for i := 0; i < 15; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)

	sum = 1

	for sum < 1000  {
		sum += sum
	}

	fmt.Println(sum)

	for {

		fmt.Println("hah")
		time.Sleep(1)
	}

	
}
