package main


func Pic(dx, dy int) [][]uint8 {

	outSlice := make([][]uint8,dy)

	for i := 0; i < len(outSlice);i++  {
		innerSlice := make([]uint8,dx)

		for j:= range innerSlice {
			innerSlice[j] = uint8(i*j - 1)
		}

		outSlice[i] = innerSlice
	}

	return outSlice

}

func main() {
}
