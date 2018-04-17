package main

import "fmt"

func main() {
	arr := []int {10,2,13,4,9,1}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for ; j > 0 && arr[j -1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

/**
插入排序，其核心思想即保证指针p之前元素 [0，p-1] 有序

时间复杂度 n*n
空间复杂度 1
 */

