package main

import (
	"strconv"
	"fmt"
)

var N []int = []int {7,2,0,4,9,1}

var M int = 233

var NLen int = len(N)

var MLen int = len(strconv.Itoa(M))

var minMax int = -1

func main()  {

	mathMax(0,0,0)
	fmt.Println(minMax)
}

func mathMax( m int,i int,j int) {
	if m > M || len(strconv.Itoa(m)) > MLen {
		fmt.Println("here")
		if m < minMax && m > M || minMax == -1 {
			minMax = m
		}
		return
	}
	for ; i < MLen; i++ {
		for ; j < NLen; j++ {
			if i != 0 || N[j] != 0 {
				m = m*10 + N[j]
				fmt.Println(m)
				fmt.Println(i + 1)
				fmt.Println(j)
				mathMax(m,i + 1,j)
			} else {
				continue
			}
		}
	}
	mathMax(0,0,j)
}