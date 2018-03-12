package main

import "fmt"

type Value interface {

	String () string
	Set (string) error
}

type Celsius float64

func FtoC(f float64) Celsius {
	return Celsius((f - 32) * 5 / 9)
}


type CelsiusFlag struct {
	Celsius
}

func (f *CelsiusFlag) Set(s string) error  {
	var value float64
	var unit string

	fmt.Sscan(s,"%f%s",&value,&unit)

	switch unit {
	case "c", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "f", "°f":
		f.Celsius = FtoC(value)
		return nil
	}
	return fmt.Errorf("invaild param %q",s)
	

}


