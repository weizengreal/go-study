package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const t = "%T(%v)\n"

	var (
		i int
		b bool
		s string
		f float64
	)

	fmt.Printf(t,ToBe,ToBe)
	fmt.Printf(t,MaxInt,MaxInt)
	fmt.Printf(t,z,z)

	fmt.Println(i,b,s,f)

	//var x , y int = 3 ,4
	//
	//var f = float64(x*x + y*y)
	//
	//var z uint = uint(f)
	//
	//fmt.Println(x,y)
	//
	//fmt.Println(f)
	//
	//fmt.Println(z)

}