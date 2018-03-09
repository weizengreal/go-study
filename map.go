package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m map[string] vertex

var newMap = map[string] vertex {
	"testIndex1" : vertex{40.321321,-70.231123},
	"testIndex2" : vertex{99.2321,-66.432},
	"testIndex3" : {22.2,33.3},
	"testIndex4" : {321.2,-998.3},
}


func main() {

	m = make(map[string]vertex)

	m["hello world index"] = vertex{1.111,2.222}

	fmt.Println(m["hello world index"])

	fmt.Println(newMap)

}

