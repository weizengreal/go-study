package main

import "fmt"

const (
	big = 1 << 100
	small = big>> 101
)


func needInt(x int) int {
	return x*10 + 1
}

func needFloat(float float64) float64 {
	return float*0.1
}

func main() {

	//fmt.Println(needInt(big))
	fmt.Println(needFloat(big))
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))



}

