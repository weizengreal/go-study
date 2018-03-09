package main

import "fmt"

type verc struct {
	x int
	y int
}


var (
	v1 = verc{1,2}
	v2 = verc{x : 1}
	v3 = verc{}
	v4 = &verc{1,2}

)




func main() {

	fmt.Println(verc{1,2})

	v := verc{3,4}

	fmt.Println(v.x)

	fmt.Println(v.y)

	v.x = 9;

	fmt.Println(v.x)

	p := &v

	fmt.Println(p.x)

	fmt.Println(p.y)

	fmt.Println(v1,v2,v3,v4,*v4)

}


