package main

import (
	"runtime"
	"fmt"
)

func main() {
	fmt.Print("now runtime is :")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("Mac os")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s",os)

	}


}


