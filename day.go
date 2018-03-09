package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("when is Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("In tow days")
	default:
		fmt.Println("long long away!")
	}


	t := time.Now()

	switch  {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	case t.Hour() < 22:
		fmt.Println("evening")
	default:
		fmt.Println("wonderful time to bed!")

	}




}

