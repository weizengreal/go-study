package main

import (
	"math"
	"fmt"
)

func main() {


	funcVar := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(funcVar(3,4))



}


