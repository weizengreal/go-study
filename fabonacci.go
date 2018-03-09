package main

import "fmt"

func fabonacci() func() int  {

	f0 , f1 , f2 := 0,1,0
	index := 0

	return func() int {
		if index == 0 {
			index ++
			return f0
		} else if index == 1 {
			index ++
			return f1
		} else {
			index ++
			f2 = f0 + f1
			f0 = f1
			f1 = f2
			return f2
		}
	}
}

func main() {
	fabo := fabonacci()

	for i := 1; i < 10; i++ {
		fmt.Println(fabo())
	}

}

