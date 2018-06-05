package main

import "fmt"

func main() {
	len := 8
	fmt.Println(fabonc(len))
	fmt.Println(fabonc_second(len))
}

func fabonc(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fabonc(n - 1) + fabonc( n - 2)
	}
}

func fabonc_second(n int) int  {
	if n == 1 || n == 2 {
		return 1
	}
	f0 := 1
	f1 := 1
	f2 := 0
	for i := 2; i < n; i++ {
		f2 = f1 + f0
		f0 = f1
		f1 = f2
	}
	return f2
}