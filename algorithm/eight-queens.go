package algorithm

import (
	"math"
	"fmt"
)

const N  = 8

var position [N]int

var countNums int = 0

func main() {
	trail(0)
}

func trail(row int) {
	if row == N {
		countNums ++
		printEq()
		return 
	}
	for column := 0; column < N; column++ {
		position[row] = column
		if vaild(row) {
			trail(row + 1)
		}
	}
}

func vaild(row int) bool {
	for i := 0;i < row; i++ {
		if position[row] == position[i] || math.Abs(float64(row - i)) == math.Abs(float64(position[row] - position[i])) {
			return false
		}
	}
	return true
}

func printEq() {
	fmt.Println(countNums)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if position[i] == j {
				fmt.Print("⊙ ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}