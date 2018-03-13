package main

import "fmt"

func main() {
	
	var n int = 6
	var start int = 64

	for i := 1; i <= n; i += 1{
		for j := 1; j <= n - i; j += 1 {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j += 1 {
			fmt.Printf("%s ",string(start + i))
		}
		fmt.Println()
	}
	
	
}
