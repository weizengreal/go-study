package main

import (
	"fmt"
)

func main()  {
	var args = make(map[string]string)

	args["test"] = "2";



	s := "test"


	fmt.Println(s)

	xByte := []rune(s)

	xByte[0] = '我'

	s = string(xByte)

	fmt.Println(s)

}