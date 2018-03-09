package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["answer"] = 42

	fmt.Println("answer:",m["answer"])

	m["answer"] = 49

	fmt.Println("answer:",m["answer"])

	delete(m,"answer")

	fmt.Println("answer:",m["answer"])

	i,v := m["answer"]

	fmt.Printf("%d -- v = %v",i,v)

}


