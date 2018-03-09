package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v Vertex) scale(f float64) {
	v.Y = v.Y * f
	v.X = v.X * f
}





type _Float64 float64

func (f _Float64) _Abs() float64 {

	if f < 20 {
		return float64(f)
	}
	return float64(20)

}


func main()  {
	v := &Vertex{25,25}
	fmt.Println(v.Abs())


	f := _Float64(-math.Sqrt2)
	fmt.Println(f._Abs())


	v1 := &Vertex{3 , 4}
	v1.scale(5)
	fmt.Println(v1.Abs())


}









