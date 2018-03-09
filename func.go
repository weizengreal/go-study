package main

import "fmt"

var python , c ,java = "python",true,1

func swap(x,y string) (string , string)  {
	return y , x
}

func split(sum int)(x , y int)  {
	x = sum * 5 / 9
	y = sum - x;
	return
}

func main() {
	var i int
	a , b  := swap("world","hello")
	fmt.Println(a,b)

	fmt.Println(split(18))

	fmt.Println(i,python,c,java)
}





