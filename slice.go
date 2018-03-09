package main

import "fmt"

func main() {

	var a []int

	printSlice("a",a)

	a = append(a,0)

	printSlice("a",a)

	a = append(a,1)

	printSlice("a",a)

	a = append(a,2,3,4,5,9)

	printSlice("a",a)


	for i,v := range a {
		fmt.Printf("%d value=%v \n",i,v)
	}

	x := make([]int ,10)

	for j:=range x{
		x[j] = 1 << uint(j)
	}

	for _,k := range x {
		fmt.Printf("value = %v \n",k)
	}




	// 2
	//a := make([]int , 0 ,5)
	//
	//fmt.Println(a,len(a),cap(a))
	//
	//if a == nil {
	//	fmt.Println("slice a is empty!")
	//} else {
	//	fmt.Println("slice a is not empty!")
	//}
	//
	//a = a[:cap(a)]
	//
	//fmt.Println(a,len(a),cap(a))
	//
	//var b []int
	//
	//fmt.Println(b,len(b),cap(b))
	//
	//if b == nil {
	//	fmt.Println("slice b is empty!")
	//}



	// 1
	//p := []int{1,2,3,4,5,6,7,8,9,10}
	//
	//fmt.Println("p=",p)
	//
	//for i := 0; i<len(p); i++  {
	//	fmt.Printf("p[%d]=%d \n",i,p[i])
	//}
	//
	//fmt.Println(p[:3])
	//
	//fmt.Println(p[3:])
	//
	//fmt.Println(p[2:5])

}

func printSlice(s string, a []int) {
	fmt.Printf("%s len:%d cap:%d %v \n",s,len(a),cap(a),a)
}
