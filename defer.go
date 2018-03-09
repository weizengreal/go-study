package main

import "fmt"

func main() {
	defer fmt.Printf(" world")
	fmt.Printf("hello ~")


	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println("hello ")
	}
	fmt.Println("ending")


}
