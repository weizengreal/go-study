package main

import "fmt"

func main() {
	arr := []int {10,2,13,4,9,1}
	quickSort(arr,0,len(arr) - 1)
	fmt.Println(arr)
}

func quickSort(arr []int, top int, rear int) {
	if top < rear {
		par := partition(arr,top,rear)
		quickSort(arr,top,par - 1)
		quickSort(arr,par + 1,rear)
	}
}

func partition(arr []int, top int, rear int) int {
	tmp := arr[rear]
	// index 指向最后一个比 tmp 小的元素
	index := top - 1

	for j := top;j < rear ; j++ {
		if arr[j] <= tmp {
			index += 1
			arr[index],arr[j] = arr[j],arr[index]
		}
	}

	// 将 partition 对应的值调整到 index 后面一个，保证 partition 之前的值小于 tmp ，之后的值大约 tmp
	arr[index + 1],arr[rear] = arr[rear],arr[index + 1]
	return index + 1
}

