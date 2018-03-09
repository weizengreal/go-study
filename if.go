package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string  {

	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}

func power(m, n, lim float64) float64 {
	if v := math.Pow(m,n);v < lim {
		return v
	}
	return lim
}



func main()  {
	fmt.Println(sqrt(2),sqrt(-5))

	fmt.Println(
		power(3,2,10),
		power(5,7,100),
	)
}
